package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/FrancescoIlario/beershop/internal/http/rest"
	bsql "github.com/FrancescoIlario/beershop/internal/storage/sql"
	_ "github.com/lib/pq"
)

const (
	EnvAddr       = "ADDR"
	EnvDbUser     = "DB_USER"
	EnvDbPassword = "DB_PASSWORD"
	EnvDbHost     = "DB_HOST"
	EnvDbPort     = "DB_PORT"
	EnvDbRetries  = "DB_RETRIES"
	EnvDbTimeout  = "DB_TIMEOUT"

	DefaultAddr       = ":8080"
	DefaultDbUser     = "postgres"
	DefaultDbPassword = "supersecret"
	DefaultDbHost     = "localhost"
	DefaultDbPort     = "5432"
	DefaultDbRetries  = 5
	DefaultDbTimeout  = 2000
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error executing server: %v", err)
	}
}

func run() error {
	fmt.Println("connecting to database")
	dbr, err := dbRetries()
	if err != nil {
		return err
	}
	dbt, err := dbTimeout()
	if err != nil {
		return err
	}

	addr := address()
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/?sslmode=disable", dbUser(), dbPassword(), dbHost(), dbPort()))
	if err != nil {
		return err
	}

	for i := 1; i <= dbr; i++ {
		if err := db.Ping(); err != nil {
			if i == dbr {
				return err
			}
			fmt.Printf("can not connect to db, retrying in %v (attempt %v/%v): %v\n", dbt, i, dbr, err)
			time.Sleep(time.Duration(dbt) * time.Millisecond)
		}
	}
	repo := bsql.New(db)

	server := rest.NewServer(repo)
	fmt.Printf("listening on %s\n", addr)
	return http.ListenAndServe(addr, server)
}

func address() string {
	return fromEnvOrDefault(EnvAddr, DefaultAddr)
}
func dbUser() string {
	return fromEnvOrDefault(EnvDbUser, DefaultDbUser)
}
func dbPassword() string {
	return fromEnvOrDefault(EnvDbPassword, DefaultDbPassword)
}
func dbHost() string {
	return fromEnvOrDefault(EnvDbHost, DefaultDbHost)
}
func dbPort() string {
	return fromEnvOrDefault(EnvDbPort, DefaultDbPort)
}
func dbRetries() (int, error) {
	return fromEnvOrDefaultInt(EnvDbRetries, DefaultDbRetries)
}
func dbTimeout() (int, error) {
	return fromEnvOrDefaultInt(EnvDbTimeout, DefaultDbTimeout)
}

func fromEnvOrDefault(env, def string) string {
	if p := os.Getenv(env); p != "" {
		return p
	}
	return def
}
func fromEnvOrDefaultInt(env string, def int) (int, error) {
	if p := os.Getenv(env); p != "" {
		return strconv.Atoi(p)
	}
	return def, nil
}
