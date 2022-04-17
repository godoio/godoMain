package databaseService

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	databaseUtils "github.com/godoio/godoMain/databaseUtils"
	task "github.com/godoio/godoMain/taskSchema"
)

func DatabaseConnect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "taskdatabase",
	}

	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return db, pingErr
	}
	fmt.Println("Connected!")
	return db, pingErr

}

func AddNewTask(newTask task.Task, db *sql.DB) (task.Task, error) {
	taskId, err := databaseUtils.CreateTask(newTask, db)
	newTask.SetId(taskId)
	if err != nil {
		log.Fatal(err)
	}
	return newTask, nil
}

func ReadExistingTasks(db *sql.DB) ([]task.Task, error) {
	tasksExisting, err := databaseUtils.ReadTasks(db)
	if err != nil {
		return tasksExisting, fmt.Errorf("read tasks: %v", err)
	}
	return tasksExisting, nil
}

func ReadExistingTask(taskId int64, db *sql.DB) (task.Task, error) {
	taskExisting, err := databaseUtils.ReadTask(taskId, db)
	if err != nil {
		return taskExisting, fmt.Errorf("read task: %v", err)
	}
	return taskExisting, nil
}

func UpdateExistingTask(taskExisting task.Task, db *sql.DB) (task.Task, error) {
	taskExisting, err := databaseUtils.UpdateTask(taskExisting, db)
	if err != nil {
		return taskExisting, fmt.Errorf("update task: %v", err)
	}
	return taskExisting, nil
}

func DeleteTask(taskId int64, db *sql.DB) error {
	err := databaseUtils.DeleteTask(taskId, db)
	if err != nil {
		return fmt.Errorf("delete task: %v", err)
	}
	return nil
}
