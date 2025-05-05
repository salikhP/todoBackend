package database

import (
	"database/sql"
)

type TaskModel struct {
	DB *sql.DB
}

type Task struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=10"`
	Date        string `json:"date" binding:"required,datetime=2006-01-02"`
}
