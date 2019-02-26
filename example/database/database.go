package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Employee struct {
	name string `json:"name"`
	city string `json:"city"`
}

func main() {
	db, err := sql.Open("mysql", "root:1@tcp(127.0.0.1:3306)/goblog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert data
	//empl := Employee{"ABC", "HCM"}
	//insertEmpl(db, empl)

	// delete data
	//deleteEmpl(db, "ABC")

	// update data
	updateEmpl(db, "huy", "bac giang")

	// select
	selectAll(db)

}

func insertEmpl(db *sql.DB, employee Employee) {
	insert, err := db.Query("INSERT INTO employee(name, city) VALUES ('" + employee.name + "','" + employee.city + "')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func deleteEmpl(db *sql.DB, emplName string) {
	del, err := db.Query("DELETE from employee WHERE name LIKE '%" + emplName + "'")
	if err != nil {
		panic(err.Error())
	}
	defer del.Close()
}

func updateEmpl(db *sql.DB, emplName string, newCity string) {
	upd, err := db.Query("UPDATE employee SET city = '" + newCity + "' WHERE name LIKE '%" + emplName + "'")
	if err != nil {
		panic(err.Error())
	}
	defer upd.Close()
}

func selectAll(db *sql.DB) {
	results, err := db.Query("SELECT name, city FROM employee")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var employee Employee
		err = results.Scan(&employee.name, &employee.city)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(employee.name + "\t\t\t" + employee.city)
	}
	//defer results.Close()
}
