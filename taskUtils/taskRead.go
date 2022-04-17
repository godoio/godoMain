package taskUtils

import (
	"database/sql"

	databaseService "github.com/godoio/godoMain/databaseService"
	task "github.com/godoio/godoMain/taskSchema"
)

func ReadTask(taskId int64) (task.Task, error) {
	var db *sql.DB
	db, err := databaseService.DatabaseConnect()
	if err != nil {
		var taskExists task.Task
		return taskExists, err
	}

	taskExists, err := databaseService.ReadExistingTask(taskId, db)
	return taskExists, err
}

func ReadTasks() ([]task.Task, error) {
	var db *sql.DB
	db, err := databaseService.DatabaseConnect()
	if err != nil {
		var taskExists []task.Task
		return taskExists, err
	}

	taskExists, err := databaseService.ReadExistingTasks(db)
	return taskExists, err
}