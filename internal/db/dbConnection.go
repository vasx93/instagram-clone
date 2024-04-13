package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DBConnection() *sql.DB {
	// func DBConnection() {

	// var uri string

	uri := os.Getenv("DB_URI")

	if uri == "" {
		log.Fatal("Invalid db uri")
	}

	fmt.Println(uri)

	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}
