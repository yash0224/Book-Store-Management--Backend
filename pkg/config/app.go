package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var d *sql.DB

// Connect initializes the database connection
func Connect() {
	// Database connection details
	dsn := "sql12729210:RxcyKMxeuS@tcp(sql12.freesqldatabase.com:3306)/sql12729210"

	// Open a connection to the database
	var err error
	d, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure it's reachable
	if err := d.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	if d == nil {
		log.Fatal("Database connection has not been established. Call Connect() first.")
	}
	return d
}
