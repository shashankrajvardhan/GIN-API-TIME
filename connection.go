package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connection() *sql.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Connected Successfully")
	// Create table if not exists
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS flight (
		flight_id SERIAL PRIMARY KEY,
		flight_name VARCHAR(255) NOT NULL,
		flight_arriving_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		flight_departure_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		passenger_details VARCHAR(255) NOT NULL,
		started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		reached_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		is_active BOOLEAN NOT NULL);
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("Created Successfully")
	return db
}
