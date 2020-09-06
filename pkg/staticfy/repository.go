package staticfy

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/prongbang/filex"
	"github.com/prongbang/staticfy/pkg/core"
)

type Repository interface {
	CreateFile(asset Assets) (Assets, error)
	DeleteFile(asset Assets) (Assets, error)
}

type repository struct {
	FileX       filex.FileX
	PqSource    DataSource
	MySQLSource DataSource
}

func (r *repository) DeleteFile(asset Assets) (Assets, error) {
	if core.IsPostgresQL() {
		if a, err := r.PqSource.Find(asset); err == nil {
			if a.ID > 0 {
				_, _ = r.FileX.Delete(a.FilePath)
				return a, r.PqSource.Delete(a)
			}
			return Assets{}, nil
		}
	}

	if a, err := r.MySQLSource.Find(asset); err == nil && a.ID > 0 {
		if a.ID > 0 {
			_, _ = r.FileX.Delete(a.FilePath)
			return a, r.MySQLSource.Delete(a)
		}
		return Assets{}, nil
	}

	return Assets{}, fmt.Errorf("%s", "Cannot delete a file")
}

func (r *repository) CreateFile(asset Assets) (Assets, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return Assets{}, err
	}

	filename := uid.String()
	if _, err = r.FileX.CreateMultipart(asset.Directory, filename, asset.FileHeader); err != nil {
		return Assets{}, err
	}
	asset.URL = asset.Host + asset.Path + "/" + filename + asset.Ext
	asset.FilePath = asset.Directory + "/" + filename + asset.Ext

	if core.IsPostgresQL() {
		return r.PqSource.Create(asset)
	}
	return r.MySQLSource.Create(asset)
}

func NewRepository(fileX filex.FileX, dataSource DataSource, mySQLSource DataSource) Repository {
	return &repository{
		FileX:       fileX,
		PqSource:    dataSource,
		MySQLSource: mySQLSource,
	}
}
