package router

import (
	"fmt"
	"net/http"

	"../routes"
	"github.com/gorilla/mux"
)

//IndexRoute export
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my API :)")
}

//Router export
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexRoute)
	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", routes.UpdateTask).Methods("PUT")

	return router
}
