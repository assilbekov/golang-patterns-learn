package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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

	// Create HTTP server
	mux := http.NewServeMux()

	// Add routes
	mux.HandleFunc("/tracker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the tracker service!")
	})

	// Start server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
