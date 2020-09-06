package main

import (
	"github.com/prongbang/staticfy/cmd"
	"os"
)

func main() {

	// TODO test only
	_ = os.Setenv("HOST", "http://localhost:4000")
	_ = os.Setenv("X-API-KEY", "df2ec6bd-eaf4-47cd-a397-1fbf1b1d60c7")
	_ = os.Setenv("JWT_SECRET", "e1305be6-0ade-4e3e-9bd2-53d15fb03d58")

	// PostgresQL
	//_ = os.Setenv("DATABASE_DRIVER_NAME", "postgres")
	//_ = os.Setenv("DATABASE_CONNECTION_STRING", "postgres://root:admin@localhost/gogs?sslmode=disable")

	// MySQL
	_ = os.Setenv("DATABASE_DRIVER_NAME", "mysql")
	_ = os.Setenv("DATABASE_CONNECTION_STRING", "root:root@(localhost:3306)/mysqlDb?charset=utf8&parseTime=true")

	cmd.Run()
}
