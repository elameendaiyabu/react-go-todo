package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	user     = "postgres"
	password = "postgres"
	port     = 5432
	host     = "localhost"
	dbname   = "todoApp"
	sslmode  = "disable"
)

func Connect() {
	connStr := fmt.Sprintf("user=%s password=%s port=%d host=%s dbname=%s sslmode=%s", user, password, port, host, dbname, sslmode)

	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := d.Ping(); err != nil {
		log.Fatal(err)
	}

	db = d

	fmt.Printf("database connection successful \n")
}

func GetDB() *sql.DB {
	return db
}
