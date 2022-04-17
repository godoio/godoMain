package databaseUtils

import (
	"database/sql"
	"fmt"
	
    task "github.com/godoio/godoMain/taskSchema"
)

func CreateTask(newtask task.Task, db *sql.DB) (int64, error) {
    result, err := db.Exec("INSERT INTO tasks (task_title, task_description, task_date) VALUES (?, ?, ?)", newtask.TaskTitle, newtask.TaskDescription, newtask.TaskDate)
    if err != nil {
        return 0, fmt.Errorf("Add task: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("Add task: %v", err)
    }
    return id, nil
}