package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", `user="hr" password="your_password" connectString="LocationOfOracleDBServer:1521/XEPDB1"`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	helloWorld(db)

	employeesNames(db)
}

// this function executes a hello world type query
func helloWorld(db *sql.DB) {
	row, err := db.Query("SELECT 'Hello World' FROM dual")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var message string

		row.Scan(&message)

		fmt.Println(message)
	}
}

// executes a SQL Query which returns the employees first and last names
func employeesNames(db *sql.DB) {
	query := "SELECT first_name, last_name FROM hr.employees ORDER BY last_name"

	employees, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer employees.Close()

	for employees.Next() {

		var firstName string
		var lastName string

		err := employees.Scan(&firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(firstName, lastName)
	}
}
