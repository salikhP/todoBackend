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
	Id            int    `json:"id"`
	UserId        int    `json:"userId" binding:"required"`
	TypeId        int    `json:"typeId" binding:"required"`
	Title         string `json:"title" binding:"required,min=3"`
	Description   string `json:"description" binding:"required,min=10"`
	Date          string `json:"date" binding:"required,datetime=2006-01-02 15:04:05"`
	TotalUnits    int    `json:"totalUnits,omitempty"`
	ProgressUnits int    `json:"progressUnits,omitempty"`
}

func (m TaskModel) Insert(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO tasks (user_id, type_id, title, description, date, total_units, progress_units) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := m.DB.QueryRowContext(ctx, query,
		task.UserId,
		task.TypeId,
		task.Title,
		task.Description,
		task.Date,
		task.TotalUnits,
		task.ProgressUnits,
	).Scan(&task.Id)

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
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.TypeId,
			&task.Title,
			&task.Description,
			&task.Date,
			&task.TotalUnits,
			&task.ProgressUnits,
		)
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

	err := row.Scan(
		&task.Id,
		&task.UserId,
		&task.TypeId,
		&task.Title,
		&task.Description,
		&task.Date,
		&task.TotalUnits,
		&task.ProgressUnits,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (m TaskModel) Update(task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE tasks SET
                 type_id = $1,
                 title = $2, 
                 description = $3, 
                 date = $4,
                 total_units = $5,
                 progress_units = $6
             WHERE id = $7`

	_, err := m.DB.ExecContext(ctx, query,
		task.TypeId,
		task.Title,
		task.Description,
		task.Date,
		task.TotalUnits,
		task.ProgressUnits,
		task.Id,
	)

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

func (m TaskModel) GetTasksByUserId(userId int) ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM tasks WHERE user_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.TypeId,
			&task.Title,
			&task.Description,
			&task.Date,
			&task.TotalUnits,
			&task.ProgressUnits,
		)
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
