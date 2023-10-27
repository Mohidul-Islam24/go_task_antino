package db

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var Driver *sql.DB
var once sync.Once

func init() {
	// get database URL
	// for testing please create a new database and update dbURL
	dbURL := "postgres://postgres@localhost:5432/antinotestdb?sslmode=disable"

	// open the database
	pgDriver, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err.Error())
	}

	pgDriver.SetMaxOpenConns(2)
	pgDriver.SetMaxIdleConns(1)
	pgDriver.SetConnMaxLifetime(time.Hour)

	Driver = pgDriver
}
