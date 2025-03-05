package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func main() {
	db, err := sql.Open("postgres", "postgres://root:qwerty@localhost:5433/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database connection established")
}
