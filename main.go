package main

import "database/sql"

var db *sql.DB
var flight []Flight

func main() {
	connection()
	router()
}
