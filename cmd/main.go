package main

import (
	"go_task_list/internal/pkg/db"
	"go_task_list/web"
	"log"
)

func main() {
	err, db := db.InitDB()
	if err != nil {
		db.Close()
		log.Fatal(err.Error() + "error during db init")
	} else {
		defer db.Close()
	}
	web.StartRouter(db)
}
