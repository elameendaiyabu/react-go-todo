package controllers

import (
	"fmt"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting todos \n")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Creating Todo \n")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating Todo \n")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Deleting todo \n")
}
