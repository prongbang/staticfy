package staticfy

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type pqDataSource struct {
	SqlX *sqlx.DB
}

func (d *pqDataSource) Find(asset Assets) (Assets, error) {
	data := Assets{}
	err := d.SqlX.Get(&data, "SELECT * FROM assets WHERE id=$1 AND user_id=$2", asset.ID, asset.UserID)
	if err == nil {
		return data, nil
	}
	return Assets{}, err
}

func (d *pqDataSource) Create(asset Assets) (Assets, error) {
	sql := `INSERT INTO assets (url, user_id, file_path) VALUES (:url, :userId, :filePath) RETURNING id`
	args := map[string]interface{}{
		"url":      asset.URL,
		"userId":   asset.UserID,
		"filePath": asset.FilePath,
	}
	stmt, err := d.SqlX.PrepareNamed(sql)
	if err == nil {
		var id int64
		err := stmt.Get(&id, args)
		if err == nil {
			asset.ID = id
			asset.CreatedAt = time.Now()
			asset.UpdatedAt = time.Now()
			return asset, nil
		}
	}
	log.Println(err)
	return Assets{}, fmt.Errorf("%s", "Cannot add assets")
}

func (d *pqDataSource) Delete(asset Assets) error {
	rs, err := d.SqlX.Exec("DELETE FROM assets WHERE id=$1 AND user_id=$2", asset.ID, asset.UserID)
	if err == nil {
		if id, err := rs.RowsAffected(); err == nil && id > 0 {
			return nil
		}
	}
	return fmt.Errorf("%s", "Cannot delete a child row")
}

// NewPqDataSource is instance
func NewPqDataSource(sqlX *sqlx.DB) DataSource {
	return &pqDataSource{
		SqlX: sqlX,
	}
}
