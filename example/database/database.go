package main

import "database/sql"

type Employee struct {
	id	int
	name	string
	address	string
}

func insertData(db *sql.DB, employee Employee)  {
	insert, _ := db.Query("INSERT INTO employee VALUES ("+employee.name+","+employee.address+")")
	defer insert.Close()
}

func main() {
	db, err := sql.Open("mysql", "huynq:1@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	// insert data
}


