package taskService

import (
	task "github.com/godoio/godoMain/taskSchema"
	taskUtils "github.com/godoio/godoMain/taskUtils"
)

func CreateTask(taskTitle, taskDescription string) (task.Task, error) {
	newTask, err := taskUtils.CreateTask(taskTitle, taskDescription)
	return newTask, err
}

func ReadTask(taskId int64) (task.Task, error) {
	taskExists, err := taskUtils.ReadTask(taskId)
	return taskExists, err
}

func ReadTasks() ([]task.Task, error) {
	tasksExist, err := taskUtils.ReadTasks()
	return tasksExist, err
}

func UpdateTask(taskTitle, taskDescription string, taskId int64) (task.Task, error) {
	taskExists, err := taskUtils.UpdateTask(taskTitle, taskDescription, taskId)
	return taskExists, err
}

func DeleteTask(taskId int64) error {
	err := taskUtils.DeleteTask(taskId)
	return err
}
