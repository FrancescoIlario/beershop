package sql_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/FrancescoIlario/beershop"
	bsql "github.com/FrancescoIlario/beershop/internal/storage/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/matryer/is"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	database := "beers"
	resource, err := pool.Run("postgres", "12.4", []string{"POSTGRES_PASSWORD=supersecret", "POSTGRES_DB=" + database})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	port := resource.GetPort("5432/tcp")
	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:supersecret@localhost:%s/%s?sslmode=disable", port, database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	c := `CREATE TABLE beers(id uuid PRIMARY KEY, name varchar(100), abv FLOAT)`
	if _, err := db.Exec(c); err != nil {
		log.Fatalf("error initialing database: %v", err)
	}

	ecode := m.Run()
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("error purging resources: %v", err)
	}
	os.Exit(ecode)
}

func Test_Create(t *testing.T) {
	is := is.New(t)
	s := bsql.New(db)
	id, err := s.Create(context.TODO(), beershop.Beer{
		Abv:  1.0,
		Name: "beer 1",
	})
	is.NoErr(err)
	is.True(id != uuid.Nil)
}

func Test_Read(t *testing.T) {
	is := is.New(t)
	ctx := context.TODO()
	s := bsql.New(db)
	b := beershop.Beer{
		Abv:  1.0,
		Name: "beer 1",
	}
	id, err := s.Create(ctx, b)
	is.NoErr(err)
	is.True(id != uuid.Nil)
	b.ID = id

	br, err := s.Read(ctx, id)
	is.NoErr(err)
	is.Equal(b, br)
}

func Test_Delete(t *testing.T) {
	is := is.New(t)
	ctx := context.TODO()
	s := bsql.New(db)
	b := beershop.Beer{
		Abv:  1.0,
		Name: "beer 1",
	}
	id, err := s.Create(ctx, b)
	is.NoErr(err)
	is.True(id != uuid.Nil)

	err = s.Delete(ctx, id)
	is.NoErr(err)
}

func Test_List_NoIsolation(t *testing.T) {
	is := is.New(t)
	ctx := context.TODO()
	s := bsql.New(db)
	bb := []beershop.Beer{
		{Abv: 1.0, Name: "beer 1"},
		{Abv: 2.0, Name: "beer 2"},
	}
	for _, b := range bb {
		id, err := s.Create(ctx, b)
		is.NoErr(err)
		is.True(id != uuid.Nil)
		b.ID = id
	}

	bbl, err := s.List(ctx)
	is.NoErr(err)
	is.True(len(bbl) >= 2)
	is.True(bbl[0].ID != uuid.Nil)
}
