package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../db" //tasks
	"github.com/gorilla/mux"
)

//GetTasks exported
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Tasks)
}

//CreateTask exported
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var NewTask db.Task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task")
	}
	fmt.Println(reqBody)
	json.Unmarshal(reqBody, &NewTask)

	NewTask.ID = len(db.Tasks) + 1

	db.Tasks = append(db.Tasks, NewTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewTask)
}

//GetTask exported
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //extrae variables desde el request
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for _, task := range db.Tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

//DeleteTask exported
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for index, task := range db.Tasks {
		if task.ID == taskID {
			db.Tasks = append(db.Tasks[:index], db.Tasks[index+1:]...)
			//:index conservar todo loq ue este antes del indice y concatenalo con lo que esta despues
			fmt.Fprintf(w, "The task with ID %v has been deleted successfully", taskID)
		}
	}

}

//UpdateTask exported
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask db.Task

	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}
	//ioutil lee body
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "please enter valid data")
	}
	//asignar reqbody a updatedtask
	json.Unmarshal(reqbody, &updatedTask)

	for index, task := range db.Tasks {
		if task.ID == taskID {
			db.Tasks = append(db.Tasks[:index], db.Tasks[index+1:]...)
			updatedTask.ID = taskID
			db.Tasks = append(db.Tasks, updatedTask)

			fmt.Fprintf(w, "Tasks with ID %v has been updated successfully", taskID)
		}
	}
}
