package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/FrancescoIlario/beershop/internal/beershop/http/rest"
	"github.com/FrancescoIlario/beershop/internal/storage/inmem"
)

const addr = ":8080"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error executing server: %v", err)
	}
}

func run() error {
	fmt.Println("running server")
	db := inmem.New()
	server := rest.NewServer(db)

	fmt.Printf("listening on %s", addr)
	err := http.ListenAndServe(addr, server)
	return err
}
