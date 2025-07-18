package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5"
)

type Connector interface {
	Connect()
}

type PostgresSource struct {
	conn *sql.DB
}

func NewPostgresSource(postgresConnector *sql.DB) (PostgresSource, error) {
	return PostgresSource{
		conn: postgresConnector,
	}, nil

}

func main() {

	dbcon, err := sql.Open("psql", "localhost:5432")
	if err != nil {
		fmt.Printf("Error opening DB connection: %v\n", err)
		os.Exit(1)
	}

	pgs, err := NewPostgresSource(dbcon)
	if err != nil {
		println("Error setting up PostgresSource: %v", err)
		os.Exit(1)
	}
	println("%v", pgs.conn)
	println("finished!!")

}
