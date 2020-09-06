package core

import (
	"os"
)

func IsPostgresQL() bool {
	return os.Getenv("DATABASE_DRIVER_NAME") == "postgres"
}

func IsMySQL() bool {
	return os.Getenv("DATABASE_DRIVER_NAME") == "mysql"
}
