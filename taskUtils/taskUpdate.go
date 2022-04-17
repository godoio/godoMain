package taskUtils

import (
	task "github.com/godoio/godoMain/taskSchema"

	databaseService "github.com/godoio/godoMain/databaseService"
)

func UpdateTask(taskTitle, taskDescription string, taskId int64) (task.Task, error) {
	taskExisting, err := ReadTask(taskId)
	if err != nil {
		return taskExisting, err
	}
	taskExisting.SetTaskTitle(taskTitle)
	taskExisting.SetTaskDescription(taskDescription)
	db, err := databaseService.DatabaseConnect()
	if err != nil {
		return taskExisting, err
	}
	taskExistingUpdate, err := databaseService.UpdateExistingTask(taskExisting, db)
	if err != nil {
		return taskExisting, err
	}
	return taskExistingUpdate, nil
}