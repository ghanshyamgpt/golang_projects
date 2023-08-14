package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	r := mux.NewRouter()

	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Content: "Do something"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Content: "Do something else"})

	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
