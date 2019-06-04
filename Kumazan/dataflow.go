package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xibinliu/dataflow/Kumazan/people"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", ":memory:")
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
			Lastname:  "Green",
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

	PPeoples := []people.PPeople{}
	err = db.Select(&PPeoples, "SELECT * FROM people")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range PPeoples {
		var pp people.Presenter = &PPeoples[i]
		pp.Present()
	}
	bytes, err := json.Marshal(PPeoples)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	http.HandleFunc("/peoples", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bytes)
	})
	http.ListenAndServe(":8080", nil)
}
