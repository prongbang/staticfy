package cmd

import (
	"github.com/prongbang/filex"
	"github.com/prongbang/staticfy/pkg/core"
	"github.com/prongbang/staticfy/pkg/database"
	"github.com/prongbang/staticfy/pkg/staticfy"
)

func Run() {
	fileX := filex.New()
	sqlX := database.NewSQLxDb()
	conn := sqlX.Connect()
	pqSource := staticfy.NewPqDataSource(conn)
	mysqlSource := staticfy.NewMySQLDataSource(conn)
	if core.IsPostgresQL() {
		_, _ = conn.Exec(staticfy.PostgresQLSchema)
	} else if core.IsMySQL() {
		_, _ = conn.Exec(staticfy.MySQLSchema)
	}
	repos := staticfy.NewRepository(fileX, pqSource, mysqlSource)
	useCase := staticfy.NewUseCase(repos)
	handle := staticfy.NewHandler(useCase)
	route := staticfy.NewRouter(handle)
	route.Register()
}
