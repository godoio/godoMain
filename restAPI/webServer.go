package restAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	task "github.com/godoio/godoMain/taskSchema"
	taskService "github.com/godoio/godoMain/taskService"
)

func viewTaskHandler(taskWriter http.ResponseWriter, request *http.Request) {
	var taskExists task.Task
	var err error
	queryParam := request.URL.Query().Get("id")
	param, err1 := strconv.ParseInt(queryParam, 10, 64)
	if err1 != nil {
		fmt.Fprintf(taskWriter, "Wrong Id")
	}
	taskExists, err = taskService.ReadTask(param)
	if err != nil {
		fmt.Fprintf(taskWriter, "Task doesnt exist")
	}
	taskJson, err2 := json.Marshal(taskExists)
	if err2 != nil {
		fmt.Fprintf(taskWriter, "%s", err2)
		return
	}
	taskWriter.Header().Set("Access-Control-Allow-Origin", "*")
	taskWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(taskWriter, "%s", string(taskJson))
}

func viewTasksHandler(taskWriter http.ResponseWriter, request *http.Request) {
	var tasks []task.Task
	var err error
	tasks, err = taskService.ReadTasks()
	if err != nil {
		fmt.Fprintf(taskWriter, "Task doesnt exist")
	}
	taskJson, err2 := json.Marshal(tasks)
	if err2 != nil {
		fmt.Fprintf(taskWriter, "%s", err2)
		return
	}
	taskWriter.Header().Set("Access-Control-Allow-Origin", "*")
	taskWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(taskWriter, "%s", string(taskJson))
}

func deleteTaskHandler(taskWriter http.ResponseWriter, request *http.Request) {
	queryParam := request.URL.Query().Get("id")
	param, err1 := strconv.ParseInt(queryParam, 10, 64)
	type delResponse struct {
		TaskId  int64
		Deleted bool
	}
	var res = delResponse{param, true}
	if err1 != nil {
		fmt.Fprintf(taskWriter, "Wrong Id")
	}
	err := taskService.DeleteTask(param)
	if err != nil {
		fmt.Fprintf(taskWriter, "Task doesnt exist")
	}
	taskJson, err2 := json.Marshal(res)
	if err2 != nil {
		fmt.Fprintf(taskWriter, "%s", err2)
		return
	}

	taskWriter.Header().Set("Access-Control-Allow-Origin", "*")
	taskWriter.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(taskWriter, "%s", string(taskJson))
}

func createTaskHandler(taskWriter http.ResponseWriter, request *http.Request) {
	queryParamTitle := request.URL.Query().Get("title")
	queryParamDescription := request.URL.Query().Get("desc")
	task, err := taskService.CreateTask(queryParamTitle, queryParamDescription)
	if err != nil {
		fmt.Fprintf(taskWriter, "Wrong input")
	}
	taskJson, err2 := json.Marshal(task)
	if err2 != nil {
		fmt.Fprintf(taskWriter, "%s", err2)
		return
	}
	taskWriter.Header().Set("Access-Control-Allow-Origin", "*")
	taskWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(taskWriter, "%s", string(taskJson))
}

func editTaskHandler(taskWriter http.ResponseWriter, request *http.Request) {
	queryParamTitle := request.URL.Query().Get("title")
	queryParamDescription := request.URL.Query().Get("desc")
	queryParamId := request.URL.Query().Get("id")
	paramId, err1 := strconv.ParseInt(queryParamId, 10, 64)
	if err1 != nil {
		fmt.Fprintf(taskWriter, "Wrong Id")
	}
	task, err := taskService.UpdateTask(queryParamTitle, queryParamDescription, paramId)
	if err != nil {
		fmt.Fprintf(taskWriter, "Wrong input")
	}
	taskJson, err2 := json.Marshal(task)
	if err2 != nil {
		fmt.Fprintf(taskWriter, "%s", err2)
		return
	}
	taskWriter.Header().Set("Access-Control-Allow-Origin", "*")
	taskWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(taskWriter, "%s", string(taskJson))
}

func GetUI() {
	http.HandleFunc("/view", viewTaskHandler)
	http.HandleFunc("/delete", deleteTaskHandler)
	http.HandleFunc("/create", createTaskHandler)
	http.HandleFunc("/edit", editTaskHandler)
	http.HandleFunc("/viewAll", viewTasksHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
