package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
)

const (
	host = "localhost"
	port = 5432
	user = "kdonati"
	// password = nil
	dbname = "kdonati"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {

	ctx := context.Background()
	pgURL := fmt.Sprintf("postgres://%s@%s:%d/%s", user, host, port, dbname)
	dbpool, err := pgxpool.New(ctx, pgURL)
	if err != nil {
		log.Fatalf("there was an error: %v", err)
	}

	defer dbpool.Close()

	rows, err := dbpool.Query(ctx, "select * from users;")
	if err != nil {
		log.Fatalf("Query failed with %s", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.Age, &u.FirstName, &u.LastName, &u.Email); err != nil {
			log.Fatalf("issues getting row, %v", err)
		}

		users = append(users, u)

	}

	if err := rows.Err(); err != nil {
		log.Fatalf("bad stuff happened %v", err)
		os.Exit(1)
	}

	log.Println("--------- All db rows: -------")
	for _, user := range users {
		log.Printf("user: %v\n", user)
	}

}
