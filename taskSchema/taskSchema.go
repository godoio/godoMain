package taskSchema

import (
	"strings"
	"time"
)

type Task struct {
	TaskId int64
	TaskTitle string
	TaskDate string
	TaskDescription string
}

func (task *Task) SetId(taskId int64) {
	task.TaskId = taskId
}

func (task *Task) InitTask() {
	current_Time := time.Now()
	task.TaskDate = strings.Split(current_Time.String(), " +")[0]
}

func (task *Task) SetTaskTitle(taskTitle string) {
	task.TaskTitle = taskTitle
}

func (task *Task) SetTaskDescription(taskDescription string) {
	task.TaskDescription = taskDescription
}
