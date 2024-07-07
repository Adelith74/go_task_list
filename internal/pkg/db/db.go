package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	var err error
	connStr := "user=postgres password=1234 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS task_tracker")
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
