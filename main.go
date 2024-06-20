package main

import (
	"baverly/config"
	"baverly/db/migrations"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {

	var conf config.AppConfig
	err := config.GetConfig(&conf)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(conf.DSN, s)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "test Beverly Hills server")
	})

	e.Logger.Fatal(e.Start(":8088"))
}

func runDBMigrate(dsn string, source *bindata.AssetSource) {
	d, err := bindata.WithInstance(source)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Warn(err)
		} else {
			log.Fatal(err)
		}
	}
}
