package db

import (
	"database/sql"
	"errors"
	"go_task_list/internal/models"
	_ "go_task_list/internal/models"
	"time"
)

type DB_Helper struct {
	Db *sql.DB
}

func GetHelper() *DB_Helper {
	return &DB_Helper{}
}

func (helper *DB_Helper) GetUser(username, password string) (*models.User, error) {
	row, err := helper.Db.Query("SELECT * FROM public.users WHERE username=$1 AND password=$2", username, password)
	if err != nil {
		return nil, errors.New(err.Error() + ": error occured at db_helper.go:20")
	}
	user := &models.User{}
	row.Scan(&user.Id, &user.Username)
	return user, nil
}

func (helper *DB_Helper) GetAllTasks() ([]models.Task, error) {
	result, err := helper.Db.Query("SELECT * FROM public.tasks")
	if err != nil {
		return nil, errors.New("unable to fetch tasks from db")
	}
	tasks := []models.Task{}
	for result.Next() {
		task := models.Task{}
		err := result.Scan(&task.User_id, &task.Task_id, &task.Created_at, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (helper *DB_Helper) InsertTask(user_id int, created_at time.Time, title string, description string) error {
	_, err := helper.Db.Query("INSERT INTO public.tasks (user_id, created_at, title, description, status) VALUES($1, $2, $3, $4, false)", user_id, created_at, title, description)
	if err != nil {
		return errors.New(err.Error() + ": error at db_helper.go:47")
	}
	return nil
}
