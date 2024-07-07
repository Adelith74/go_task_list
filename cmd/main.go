package main

import (
	//"go_task_list/internal/pkg/db"
	"go_task_list/web"
)

func main() {
	//db.InitDB()
	web.StartRouter()
}
