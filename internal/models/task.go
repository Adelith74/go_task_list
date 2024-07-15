package models

import (
	"time"
)

type Task struct {
	User_id     int       `db:"user_id" json:"user_id"`
	Task_id     int       `db:"task_id" json:"task_id"`
	Created_at  time.Time `db:"created_at" json:"created_at"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      bool      `db:"status" json:"status"`
}
