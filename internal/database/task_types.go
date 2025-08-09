package database

import (
	"context"
	"database/sql"
	"time"
)

type TaskTypeModel struct {
	DB *sql.DB
}

type TaskType struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required,min=1"`
	Unit string `json:"unit" binding:"required,min=1"`
}

func (m TaskTypeModel) Insert(taskType *TaskType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO task_types(name, unit) VALUES ($1, $2) RETURNING id`

	err := m.DB.QueryRowContext(ctx, query,
		taskType.Name,
		taskType.Unit,
	).Scan(&taskType.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m TaskTypeModel) GetAll() ([]*TaskType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM task_types`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskTypes := []*TaskType{}

	for rows.Next() {
		var taskType TaskType
		err := rows.Scan(
			&taskType.Id,
			&taskType.Name,
			&taskType.Unit,
		)
		if err != nil {
			return nil, err
		}
		taskTypes = append(taskTypes, &taskType)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return taskTypes, nil
}

func (m TaskTypeModel) Get(id int) (*TaskType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM task_types WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var taskType TaskType

	err := row.Scan(
		&taskType.Id,
		&taskType.Name,
		&taskType.Unit,
	)

	if err != nil {
		return nil, err
	}

	return &taskType, nil
}

func (m TaskTypeModel) Update(taskType *TaskType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE task_types SET name = $1, unit = $2 WHERE id = $3`

	_, err := m.DB.ExecContext(ctx, query,
		taskType.Name,
		taskType.Unit,
		taskType.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m TaskTypeModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM task_types WHERE id = $1`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
