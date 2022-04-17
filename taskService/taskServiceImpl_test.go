package taskService

import (
	"fmt"
	"testing"
)

func TestCreateTask(t *testing.T) {
	taskTitle := "A beautiful day!"
	taskDescription := "A beautiful day in God's Providence!"

	newTask, err := CreateTask(taskTitle, taskDescription)

	if err != nil {
		t.Fatalf(`Error on empty title and/or description %v`, err)
	} else {
		fmt.Println("Task ID:", newTask.TaskId)
		fmt.Println("Task created:", newTask.TaskDate)
		fmt.Println("Task Title:", newTask.TaskTitle)
		fmt.Println("Task Description:", newTask.TaskDescription)
	}
}

func TestReadTask(t *testing.T) {
	var taskId int64 = 11

	newTask, err := ReadTask(taskId)

	// newTask, err2 := SaveTask(newTask)
	if err != nil {
		t.Fatalf(`Error on empty title and/or description %v`, err)
	} else {
		fmt.Println("Task ID:", newTask.TaskId)
		fmt.Println("Task created:", newTask.TaskDate)
		fmt.Println("Task Title:", newTask.TaskTitle)
		fmt.Println("Task Description:", newTask.TaskDescription)
	}
}

func TestReadTasks(t *testing.T) {

	tasks, err := ReadTasks()

	// newTask, err2 := SaveTask(newTask)
	if err != nil {
		t.Fatalf(`Error on empty title and/or description %v`, err)
	} else {
		i := 0
		for i = 0; i < len(tasks); i++ {
			fmt.Println("Task ID:", tasks[i].TaskId)
			fmt.Println("Task created:", tasks[i].TaskDate)
			fmt.Println("Task Title:", tasks[i].TaskTitle)
			fmt.Println("Task Description:", tasks[i].TaskDescription)
		}
	}
}

func TestUpdateTask(t *testing.T) {
	var taskId int64 = 11
	taskTitle := "A beautiful day!"
	taskDescription := "A beautiful day in God's Providence! In the name of God the Father, and the Son and the Holy Spirit. Amen!"

	newTask, err := UpdateTask(taskTitle, taskDescription, taskId)

	// newTask, err2 := SaveTask(newTask)
	if err != nil {
		t.Fatalf(`Error on empty title and/or description %v`, err)
	} else {
		fmt.Println("Task ID:", newTask.TaskId)
		fmt.Println("Task updated:", newTask.TaskDate)
		fmt.Println("Task Title:", newTask.TaskTitle)
		fmt.Println("Task Description:", newTask.TaskDescription)
	}
}

func TestDeleteTask(t *testing.T) {
	var taskId int64 = 10
	err := DeleteTask(taskId)
	if err != nil {
		t.Fatalf(`Error on empty title and/or description %v`, err)
	} else {
		fmt.Println("Task Deleted succesfully!")
	}
}
