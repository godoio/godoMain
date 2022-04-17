package databaseUtils

import (
	"database/sql"
	"fmt"
	"time"
	"strings"

    task "github.com/godoio/godoMain/taskSchema"
)

func UpdateTask(taskExisting task.Task, db *sql.DB) (task.Task, error) {
	current_Time := time.Now()
	result, err := db.Exec("UPDATE tasks SET task_title = ?, task_description = ?, task_date = ? WHERE task_id = ?", taskExisting.TaskTitle, taskExisting.TaskDescription, strings.Split(current_Time.String(), " +")[0], taskExisting.TaskId)
	if err != nil {
		var taskExisting task.Task
        return taskExisting, fmt.Errorf("Add task: %v", err)
    } 
	affect, err := result.RowsAffected()
	if err != nil {
		var taskExisting task.Task
        return taskExisting, fmt.Errorf("Add task: %v", err)
    }
	fmt.Println("Records affected", affect)
	taskExistingUpdate, err := ReadTask(taskExisting.TaskId, db) 
	if err != nil {
		var taskExisting task.Task
        return taskExisting, fmt.Errorf("Add task: %v", err)
    }
	return taskExistingUpdate, nil
}