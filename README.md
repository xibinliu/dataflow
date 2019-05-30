# dataflow
Go practice project to show how data flows between different layers

Currently in the dataflow.go file, a sqlite database is connected, a table is created, and some data is inserted into the table.

Tasks:

### 1. Directory Structure

Create your own directory with your user name and copy the dataflow.go to your own directory

### 2. Error handling

Currently, all DB operations are executed without error handling. Find out which function call will return error and add error handling (eg. print out the error and exit).

### 3. Organize the intial data

Create a structure "People" to hold the intial data and and use the structure member in the statement.Exec() call.

### 4. For loop

Put the structured People data into a slice of "Peoples", and use for loop to execute statement.Exec() call.

### 5. Package

Move the People structure into another go file as a new package, and import People from that package.

### 6. Query DB

Query from DB the inserted data and print out

### 7. Use "sqlx"

Replace the database/sql with package github.com/jmoiron/sqlx

### 8. Query DB to structure

use sqlx db.Select to query DB into the Peoples slice

### 9. Print with JSON format

Print out the Peoples slice to a json array in the standard output (using "encoding/json") package

### 10. JSON tags

Add json tags to the People structure, that when marshal the People structure, 
- id is omitted
- all keys are in lower case:
  eg: firstname, instead of Firstname

### 11. A new structure

Create a new structure PPeople, which includes all the fields of People, and and a new field "Name"

### 12. Add an interface 

Add interface "Present()" to the PPeople structure, which will combile the "Firstname" and "Lastname" to the "Name" field, and then print the whole structure to JSON

### 13. HTTP Interface

Using "net/http" to run an http server in dataflow.go, which can respond "pong" when getting http://localhost:8080/ping request

### 14. JSON response

Implement handler for "/peoples" request, and return the PPeople slice in JSON array.

### 15. Unit test

Implement unit tests for the code written and show the code coverage.



