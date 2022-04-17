package taskUtils

import (
	task "github.com/godoio/godoMain/taskSchema"
	"database/sql"

	databaseService "github.com/godoio/godoMain/databaseService"
)

func SaveTask(newTask task.Task) (task.Task, error) {
	var db *sql.DB
	db, err := databaseService.DatabaseConnect()
	if err != nil {
		return newTask, err
	}
	newTask1, err := databaseService.AddNewTask(newTask, db)
	return newTask1, err
}
