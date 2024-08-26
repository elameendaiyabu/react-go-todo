package controllers

import (
	"backend/pkg/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var db *sql.DB

type task struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting todos \n")

	w.Header().Set("Context-Type", "application/json")

	db = config.GetDB()

	sqlStatement := `SELECT id, task, status FROM tasks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		err := rows.Scan(&t.Id, &t.Task, &t.Status)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	res, _ := json.Marshal(tasks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Creating Todo \n")

	w.Header().Set("Content-Type", "application/json")

	now := time.Now()

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	task := r.FormValue("task")
	_ = json.NewDecoder(r.Body).Decode(&task)
	db = config.GetDB()

	sqlStatement := `INSERT INTO tasks ( task, created_at, status)
	VALUES ($1, $2, $3) `

	_, err := db.Exec(sqlStatement, task, now.Format("Mon Jan 2 15:04:05 MST 2006"), "todo")
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating Status \n")

	w.Header().Set("Content-Type", "application/json")
	db = config.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]
	currentStatus := vars["currentStatus"]
	if currentStatus == "todo" {

		sqlStatement := ` UPDATE tasks SET status = $1 WHERE id = $2`

		_, err := db.Exec(sqlStatement, "done", id)
		fmt.Printf("changed to done \n")
		if err != nil {
			panic(err)
		}
	} else if currentStatus == "done" {

		sqlStatement := ` UPDATE tasks SET status = $1 WHERE id = $2`

		_, err := db.Exec(sqlStatement, "todo", id)
		fmt.Printf("changed to todo \n")
		if err != nil {
			panic(err)
		}
	} else {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task status updated successfully"})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Updating Todo \n")

	db = config.GetDB()

	vars := mux.Vars(r)
	id := vars["id"]
	task := "go to the mart"

	sqlStatement := ` UPDATE tasks SET task = $1 WHERE id = $2`

	_, err := db.Exec(sqlStatement, task, id)
	if err != nil {
		panic(err)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Deleting todo \n")

	w.Header().Set("Content-Type", "application/json")

	db = config.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	sqlStatement := `DELETE FROM tasks WHERE id = $1`

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
