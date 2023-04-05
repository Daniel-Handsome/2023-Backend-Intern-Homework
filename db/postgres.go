package db

import (
	"database/sql"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	sqldb, err := initDatabase(*utils.App)

	if err != nil {
		log.Fatal(err)
	}
	return postgresNew(sqldb)
}

func postgresNew(db *sql.DB) *gorm.DB {
	ok := make(chan bool, 1)
	go func() {
		for {
			if err := db.Ping(); err != nil {
				continue
			}
			ok <- true
			return
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		panic(errors.New("db timeout for 10 second"))
	case <-ok:
		orm, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return orm
	}
}
