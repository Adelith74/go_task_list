package db

import (
	"database/sql"
	"errors"
	"fmt"
	"go_task_list/internal/models"
	_ "go_task_list/internal/models"
)

type DB_Helper struct {
	Db *sql.DB
}

func GetHelper() *DB_Helper {
	return &DB_Helper{}
}

func (helper *DB_Helper) GetAllTasks() ([]models.Task, error) {
	result, err := helper.Db.Query("SELECT * FROM public.tasks")
	if err != nil {
		return nil, errors.New("unable to fetch tasks from db")
	}
	tasks := []models.Task{}
	for result.Next() {
		task := models.Task{}
		err := result.Scan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (helper *DB_Helper) InsertTask() {
	task := models.Task{}
	fmt.Println(task)
}
