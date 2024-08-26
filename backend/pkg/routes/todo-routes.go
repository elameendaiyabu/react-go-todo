package routes

import (
	"backend/pkg/controllers"
	"backend/pkg/middleware"

	"github.com/gorilla/mux"
)

var TodoRoutes = func(r *mux.Router) {
	r.HandleFunc("/", middleware.EnableCORS(controllers.GetTasks)).Methods("GET")
	r.HandleFunc("/", middleware.EnableCORS(controllers.CreateTask)).Methods("POST")
	r.HandleFunc("/{id}", middleware.EnableCORS(controllers.UpdateTask)).Methods("PUT", "OPTIONS")
	r.HandleFunc("/status/{currentStatus}/{id}", middleware.EnableCORS(controllers.UpdateTaskStatus)).Methods("PUT", "OPTIONS")
	r.HandleFunc("/{id}", middleware.EnableCORS(controllers.DeleteTask)).Methods("DELETE", "OPTIONS")
}
