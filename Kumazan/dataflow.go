package main

import (
	"database/sql"
	"fmt"

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

	statement.Exec("Alice", "Wihte", 10, "Female")
	statement.Exec("Bob", "Green", 11, "Male")
	statement.Exec("Charles", "Brown", 12, "Male")
	statement.Exec("Dianna", "Black", 13, "Female")
}
