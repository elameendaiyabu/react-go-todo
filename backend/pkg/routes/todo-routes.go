package routes

import (
	"backend/pkg/controllers"

	"github.com/gorilla/mux"
)

var TodoRoutes = func(r *mux.Router) {
	r.HandleFunc("/", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteTodo).Methods("DELETE")
}
