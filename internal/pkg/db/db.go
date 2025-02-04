package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func createTasksTable() error {
	_, err := db.Exec(`CREATE TABLE tasks (
  	user_id SERIAL,  
	task_id SERIAL,
  	created_at TIMESTAMP NOT NULL, 
  	title VARCHAR NOT NULL,
  	description VARCHAR,
  	status BOOL NOT NULL,
	PRIMARY KEY(user_id, task_id));`)
	if err != nil {
		return err
	}
	return nil
}

func createUsersTable() error {
	_, err := db.Exec(`CREATE TABLE users (
		user_id SERIAL PRIMARY KEY,  
		username VARCHAR NOT NULL,
		password VARCHAR NOT NULL);`)
	if err != nil {
		return err
	}
	return nil
}

func checkTablesExists(table_name string) bool {
	var name string
	err := db.QueryRow("SELECT table_name FROM information_schema.tables WHERE table_name = $1", table_name).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			panic(err)
		}
	} else {
		return true
	}
}

func InitDB() (error, *sql.DB) {
	var err error
	connStr := "user=postgres password=root dbname=task_tracker sslmode=disable port=5432"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error() + " unable to establisk connection to db: row 14 of db.go")
	}
	tasksExists := checkTablesExists("tasks")
	usersExists := checkTablesExists("users")
	if !tasksExists {
		createTasksTable()
	}
	if !usersExists {
		createUsersTable()
	}
	return nil, db
}

func GetDB() *sql.DB {
	return db
}
