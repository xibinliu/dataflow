package main

import (
	"database/sql"
	"fmt"

	"github.com/xibinliu/dataflow/Kumazan/people"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println(err)
		return
	}
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS people (
			id INTEGER PRIMARY KEY,
			firstname TEXT,
			lastname TEXT,
			age INTEGER,
			gender TEXT CHECK( gender IN ('Male', 'Female') )
		)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec()
	statement, err = db.Prepare(`
		INSERT INTO people (firstname, lastname, age, gender)
		VALUES (?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	Peoples := []people.People{
		people.People{
			Firstname: "Alice",
			Lastname:  "White",
			Age:       10,
			Gender:    "Female",
		},
		people.People{
			Firstname: "Bob",
			Lastname:  "Bob",
			Age:       11,
			Gender:    "Male",
		},
		people.People{
			Firstname: "Charles",
			Lastname:  "Brown",
			Age:       12,
			Gender:    "Male",
		},
		people.People{
			Firstname: "Dianna",
			Lastname:  "Black",
			Age:       13,
			Gender:    "Female",
		},
	}
	for _, people := range Peoples {
		statement.Exec(people.Firstname, people.Lastname, people.Age, people.Gender)
	}
	rows, _ := db.Query("SELECT firstname, lastname, age, gender FROM people")
	for rows.Next() {
		var firstname string
		var lastname string
		var age int
		var gender string
		rows.Scan(&firstname, &lastname, &age, &gender)
		fmt.Println(firstname, lastname, age, gender)
	}
}
