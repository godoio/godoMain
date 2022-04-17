package taskUtils

import (
	"errors"
	task "github.com/godoio/godoMain/taskSchema"
)

func CreateTask(taskTitle, taskDescription string) (task.Task, error) {
	var newTask task.Task

	if taskTitle == "" || taskDescription == "" {
		return newTask, errors.New("empty title and/or description")
	}
	newTask.InitTask()
	newTask.SetTaskTitle(taskTitle)
	newTask.SetTaskDescription(taskDescription)
	newTask, err := SaveTask(newTask)

	return newTask, err
}