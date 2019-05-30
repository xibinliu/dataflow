package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")
	statement, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS people (
			id INTEGER PRIMARY KEY,
			firstname TEXT,
			lastname TEXT,
			age INTEGER,
			gender TEXT CHECK( gender IN ('Male', 'Female') )
		)`)
	statement.Exec()
	statement, _ = db.Prepare(`
		INSERT INTO people (firstname, lastname, age, gender)
		VALUES (?, ?, ?, ?)`)

	statement.Exec("Alice", "Wihte", 10, "Female")
	statement.Exec("Bob", "Green", 11, "Male")
	statement.Exec("Charles", "Brown", 12, "Male")
	statement.Exec("Dianna", "Black", 13, "Female")
}
