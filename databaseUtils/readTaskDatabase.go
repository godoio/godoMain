package databaseUtils

import (
	"database/sql"
	"fmt"

    task "github.com/godoio/godoMain/taskSchema"
)

func ReadTask(taskId int64, db *sql.DB) (task.Task, error) {
	
	var tasks []task.Task
	rows, err := db.Query("SELECT * FROM tasks WHERE task_id = ?", taskId)
	if err != nil {
		var taskExisting task.Task
        return taskExisting, fmt.Errorf("read task: %v", err)
    }
	defer rows.Close()
	for rows.Next() {
		var taskExisting task.Task
		if err := rows.Scan(&taskExisting.TaskId, &taskExisting.TaskTitle, &taskExisting.TaskDescription, &taskExisting.TaskDate); err != nil {
			return taskExisting, fmt.Errorf("albumsByArtist %q: %v", taskId, err)
		}
		tasks = append(tasks, taskExisting)
	}
	return tasks[0], nil
}

func ReadTasks(db *sql.DB) ([]task.Task, error) {
	
	var tasks []task.Task
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
        return tasks, fmt.Errorf("read task: %v", err)
    }
	defer rows.Close()
	for rows.Next() {
		var taskExisting task.Task
		if err := rows.Scan(&taskExisting.TaskId, &taskExisting.TaskTitle, &taskExisting.TaskDescription, &taskExisting.TaskDate); err != nil {
			return tasks, fmt.Errorf("no tasks")
		}
		tasks = append(tasks, taskExisting)
	}
	return tasks, nil
}