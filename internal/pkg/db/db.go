package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	var err error
	connStr := "user=postgres password=1234 dbname=task_tracker sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error() + " unable to establisk connection to db: row 14 of db.go")
	}
	defer db.Close()

	_, err = db.Exec("SELECT * from tasks")
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
