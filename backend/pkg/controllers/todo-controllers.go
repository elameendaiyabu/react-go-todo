package controllers

import (
	"fmt"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting todos \n")
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Creating Todo \n")
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating Todo \n")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Deleting todo \n")
}
