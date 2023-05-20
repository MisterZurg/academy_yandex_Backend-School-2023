package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/internal/config"
	"yandex-team.ru/bstask/internal/db/postgres"
	"yandex-team.ru/bstask/internal/handler"
	"yandex-team.ru/bstask/internal/repository"
	"yandex-team.ru/bstask/internal/service"
)

const (
	DBType = "postgres"
)

func main() {

	// CFG
	cfg, err := config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}

	pgDSN := config.GetPostgresConnectionString(DBType, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	pgDB, err := repository.NewRepository(pgDSN)
	if err != nil {
		panic("error parsing config")
	} else {
		fmt.Println("Connected tho")
	}

	postgres.InitTables(pgDB.DB)

	srv := service.NewService(pgDB)
	e := setupServer(srv)
	e.Logger.Fatal(e.Start(":8080"))

}

func setupServer(srv *service.Service) *echo.Echo {
	e := echo.New()
	handler.RegisterHandlersWithBaseURLUsingRateLimiter(e, srv, "")
	//routes.SetupRoutes(e)

	return e
}
