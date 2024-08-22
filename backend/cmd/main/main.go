package main

import (
	"backend/pkg/config"
	"backend/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	r := mux.NewRouter()

	routes.TodoRoutes(r)

	fmt.Printf("Starting Server at localhost:8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
