package models

import (
	"time"
)

type Task struct {
	Date        time.Time `db:"date"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      bool      `db:"status"`
}
