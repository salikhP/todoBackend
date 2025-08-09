package database

import "database/sql"

type Models struct {
	Users     UserModel
	Tasks     TaskModel
	TaskTypes TaskTypeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Tasks:     TaskModel{DB: db},
		TaskTypes: TaskTypeModel{DB: db},
	}
}
