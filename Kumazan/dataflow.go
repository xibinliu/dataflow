package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// People structure in db
type People struct {
	firstname string
	lastname  string
	age       int
	gender    string
}

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

	alice := People{
		firstname: "Alice",
		lastname:  "White",
		age:       10,
		gender:    "Female",
	}
	bob := People{
		firstname: "Bob",
		lastname:  "Bob",
		age:       11,
		gender:    "Male",
	}
	charles := People{
		firstname: "Charles",
		lastname:  "Brown",
		age:       12,
		gender:    "Male",
	}
	dianna := People{
		firstname: "Dianna",
		lastname:  "Black",
		age:       13,
		gender:    "Female",
	}
	statement.Exec(alice.firstname, alice.lastname, alice.age, alice.gender)
	statement.Exec(bob.firstname, bob.lastname, bob.age, bob.gender)
	statement.Exec(charles.firstname, charles.lastname, charles.age, charles.gender)
	statement.Exec(dianna.firstname, dianna.lastname, dianna.age, dianna.gender)
}
