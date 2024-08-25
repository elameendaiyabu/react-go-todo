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
	Id   int    `json:"id"`
	Task string `json:"task"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting todos \n")

	db = config.GetDB()

	sqlStatement := `SELECT id, task FROM tasks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tasks []task
	for rows.Next() {
		var t task
		err := rows.Scan(&t.Id, &t.Task)
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
	json.NewEncoder(w).Encode(tasks)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Creating Todo \n")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	now := time.Now()

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	task := r.FormValue("task")
	_ = json.NewDecoder(r.Body).Decode(&task)
	db = config.GetDB()

	sqlStatement := `INSERT INTO tasks ( task, created_at)
	VALUES ($1, $2) `

	_, err := db.Exec(sqlStatement, task, now.Format("Mon Jan 2 15:04:05 MST 2006"))
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(task)
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

	db = config.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	sqlStatement := `DELETE FROM tasks WHERE id = $1`

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}
