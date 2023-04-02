package db

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func New(db *sql.DB) *gorm.DB {
	ok := make(chan bool)
	go func() {
		var name string
		for {
			query := `SELECT current_database()`
			err := db.QueryRow(query).Scan(name)
			if name == "" || err != nil {
				continue
			}
			ok <- true
			return
		}
	}()

	select {
	case <-time.After(10 * time.Second):
	}
}
