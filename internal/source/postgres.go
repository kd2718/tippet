package source

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgresSource struct{}

func (PostgresSource) connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't connect to DB: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

}
