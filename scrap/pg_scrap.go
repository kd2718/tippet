package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "kdonati"
	// password = nil
	dbname = "kdonati"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.Exec(`CREATE if not exists TABLE users (
 	id SERIAL PRIMARY KEY,
  	age INT,
  	first_name TEXT,
  	last_name TEXT,
  	email TEXT UNIQUE NOT NULL
	);`)

	/*
		sqlStatement := `
			INSERT INTO users (age, email, first_name, last_name)
			VALUES (31, 'jon2@calhoun.io', 'Jonathan', 'Calhoun')`
		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	*/

	rows, err := db.Query(`select * from users;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var age int
		var first_name string
		var last_name string
		var email string
		err = rows.Scan(&id, &age, &first_name, &last_name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s\n", first_name)
	}
	//fmt.Printf("out:\n%v\n", out.)

	fmt.Println("Successfully connected!")
}
