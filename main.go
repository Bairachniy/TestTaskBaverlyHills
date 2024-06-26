package main

import (
	"baverly/config"
	"baverly/db/migrations"
	"baverly/service"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	var conf config.AppConfig
	err := config.GetConfig(&conf)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbConn := connectToDBOrFatal(conf.DSN)

	s := bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate(conf.DSN, s)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	service.InitService(dbConn, e)

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

func connectToDBOrFatal(dsn string) *sql.DB {
	driver := "postgres"
	dbConn, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Failed to open connection to database: ", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	return dbConn
}
