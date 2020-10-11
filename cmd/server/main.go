package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/FrancescoIlario/beershop/internal/http/rest"
	bsql "github.com/FrancescoIlario/beershop/internal/storage/sql"
	_ "github.com/lib/pq"
)

// TODO: read them from ENV
const (
	addr       = ":8080"
	dbUser     = "postgres"
	dbPassword = "supersecret"
	dbHost     = "postgres"
	dbPort     = "5432"
	dbRetries  = 5
	dbTimeout  = 2000
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error executing server: %v", err)
	}
}

func run() error {
	fmt.Println("connecting to database")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=disable", dbUser, dbPassword, dbHost, dbPort))
	if err != nil {
		return err
	}
	for i := 1; i <= dbRetries; i++ {
		if err := db.Ping(); err != nil {
			if i == dbRetries {
				return err
			}
			fmt.Printf("can not connect to db, retrying in %v (attempt %v/%v): %v\n", dbTimeout, i, dbRetries, err)
			time.Sleep(dbTimeout * time.Millisecond)
		}
	}
	repo := bsql.New(db)

	fmt.Printf("listening on %s", addr)
	server := rest.NewServer(repo)
	return http.ListenAndServe(addr, server)
}
