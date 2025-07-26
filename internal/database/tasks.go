package database

import (
	"context"
	"database/sql"
	"time"
)

type TaskModel struct {
	DB *sql.DB
}

type Task struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId" binding:"required"`
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=10"`
	Date        string `json:"date" binding:"required,datetime=2006-01-02"`
}

func (m TaskModel) Insert(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO tasks (user_id, title, description, date) VALUES ($1, $2, $3, $4) RETURNING id`

	err := m.DB.QueryRowContext(ctx, query, task.UserId, task.Title, task.Description, task.Date).Scan(&task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m TaskModel) GetAll() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM tasks`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Date)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m TaskModel) Get(id int) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM tasks WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var task Task

	err := row.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Date)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (m TaskModel) Update(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE tasks SET title = $1, description = $2, date = $3 WHERE id = $4`

	_, err := m.DB.ExecContext(ctx, query, task.Title, task.Description, task.Date, task.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m TaskModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM tasks WHERE id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
