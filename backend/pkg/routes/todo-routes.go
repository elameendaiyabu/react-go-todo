package routes

import (
	"backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var TodoRoutes = func(r *mux.Router) {
	r.HandleFunc("/", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
}
