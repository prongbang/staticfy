package staticfy

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type mySQLDataSource struct {
	SqlX *sqlx.DB
}

func (d *mySQLDataSource) Find(asset Assets) (Assets, error) {
	sql := `SELECT id, url, user_id, file_path, created_at, updated_at FROM assets WHERE id=? AND user_id=?`
	var rows []Assets
	if err := d.SqlX.Select(&rows, sql, asset.ID, asset.UserID); err != nil {
		log.Println(err)
		return Assets{}, err
	}
	if len(rows) > 0 {
		return rows[0], nil
	}
	return Assets{}, fmt.Errorf("%s", "Not found")
}

func (d *mySQLDataSource) Create(asset Assets) (Assets, error) {
	sql := `INSERT INTO assets (url, user_id, file_path) VALUES (?, ?, ?)`
	rs, err := d.SqlX.Exec(sql, asset.URL, asset.UserID, asset.FilePath)
	if err == nil {
		id, err := rs.LastInsertId()
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

func (d *mySQLDataSource) Delete(asset Assets) error {
	rs, err := d.SqlX.Exec("DELETE FROM assets WHERE id=? AND user_id=?", asset.ID, asset.UserID)
	if err == nil {
		if id, err := rs.RowsAffected(); err == nil && id > 0 {
			return nil
		}
	}
	return fmt.Errorf("%s", "Cannot delete a child row")
}

// NewMySQLDataSource is instance
func NewMySQLDataSource(sqlX *sqlx.DB) DataSource {
	return &mySQLDataSource{
		SqlX: sqlX,
	}
}
