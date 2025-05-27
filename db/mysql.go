package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"game-service/config"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", config.DSN)
		if err != nil {
			log.Fatalf("Cannot connect to MySQL: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Cannot ping MySQL: %v", err)
		}
	})
	return db
}
