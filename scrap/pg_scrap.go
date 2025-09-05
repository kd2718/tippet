package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
	//"github.com/jackc/pgx/v5/pgxpool"
	//_ "github.com/jackc/pgx/v5/pgxpool"
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
	/*
		for this i had to "brew install tomlplusplus"
		then I had to set some env variables
			export CGO_CPPFLAGS="-I$(brew --prefix tomlplusplus)/include"
			export CGO_LDFLAGS="-L$(brew --prefix tomlplusplus)/lib"

		I don't think these are needed after I build, but I am not sure
	*/
	var drv drivermgr.Driver
	ctx := context.Background()
	pgURL := fmt.Sprintf("postgres://%s@%s:%d/%s", user, host, port, dbname)
	//alloc := memory.DefaultAllocator

	// 2. Instantiate the ADBC PostgreSQL driver
	// The driver requires the memory allocator.
	//drv := postgres.NewDriver(alloc)
	//pg := postgres.New()
	db, err := drv.NewDatabase(map[string]string{
		"driver":          "abc_driver_postgres",
		adbc.OptionKeyURI: pgURL,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cxn, err := db.Open(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer cxn.Close()

	stmnt, err := cxn.NewStatement()
	if err != nil {
		log.Fatal(err)
	}

	defer stmnt.Close()

	query := "select * from users;"
	if err := stmnt.SetSqlQuery(query); err != nil {
		log.Fatal(err)
	}

	rdr, rows, err := stmnt.ExecuteQuery(ctx)
	defer rdr.Release()

	fmt.Printf("read %s rows\n", rows)

	fmt.Printf("****\n")
	rdr.Schema()

	for rdr.Next() {
		rec := rdr.Record()
		fmt.Printf("%s", rec)
	}

	fmt.Printf("The End")

}

/*
func main2() {

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

	fmt.Printf("desc: - %v", rows.FieldDescriptions())
	rows.

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
*/
