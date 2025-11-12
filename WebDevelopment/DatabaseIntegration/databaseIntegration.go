package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Nishada", 25)
	if err != nil {
		panic(err)
	}

	// Query
	rows, _ := db.Query("SELECT name, age FROM users")
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		fmt.Println(name, age)
	}
}
