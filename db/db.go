package db

import (
	"database/sql"
	"fmt"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func initDatabase(config utils.Config) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s://%s:%s@%v:%v/%s?sslmode=disable",
		config.Connection,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err = sql.Open(config.Connection, dsn)
	if err != nil {
		return
	}

	err = initMigrate(db, config.Database)

	return db, err
}

func Close() {
	db.Close()
}

func initMigrate(db *sql.DB, database string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		database,
		driver,
	)
	// or m.Step(2) if you want to explicitly set the number of migrations to run

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("---------  initializing migrations is successful ----")
	return nil
}
