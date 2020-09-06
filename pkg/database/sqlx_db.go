package database

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// SQLxDb the interface
type SQLxDb interface {
	Connect() *sqlx.DB
}

type sqlXDb struct {
}

func (s *sqlXDb) Connect() *sqlx.DB {
	db, err := sqlx.Connect(os.Getenv("DATABASE_DRIVER_NAME"), os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		log.Println(err)
		return &sqlx.DB{}
	}
	log.Println("Database " + os.Getenv("DATABASE_DRIVER_NAME") + " Connected.")
	return db
}

// NewSQLxDb is instance
func NewSQLxDb() SQLxDb {
	return &sqlXDb{}
}
