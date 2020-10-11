package sql

import (
	"context"
	"database/sql"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) beershop.Repository {
	return &store{db: db}
}

func (s *store) Create(ctx context.Context, b beershop.Beer) (uuid.UUID, error) {
	id := uuid.New()
	it := `INSERT INTO beers(id, name, abv) values ($1, $2, $3)`
	if _, err := s.db.ExecContext(ctx, it, id, b.Name, b.Abv); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *store) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var i uuid.UUID
	q := `SELECT id FROM beers WHERE id = $1`
	if err := tx.QueryRowContext(ctx, q, id).Scan(&i); err != nil {
		return err
	}

	e := `DELETE FROM beers WHERE id = $1`
	r, err := tx.ExecContext(ctx, e, id)
	if err != nil {
		return err
	}
	if _, err := r.RowsAffected(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *store) List(ctx context.Context) ([]beershop.Beer, error) {
	var bb []beershop.Beer
	q := `SELECT id, name, abv FROM beers`
	r, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var id uuid.UUID
		var name string
		var abv float32
		if err := r.Scan(&id, &name, &abv); err != nil {
			return nil, err
		}
		b := beershop.Beer{
			ID:   id,
			Name: name,
			Abv:  abv,
		}
		bb = append(bb, b)
	}
	if err := r.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}

func (s *store) Read(ctx context.Context, id uuid.UUID) (beershop.Beer, error) {
	var name string
	var abv float32
	q := `SELECT name, abv FROM beers WHERE id = $1`
	if err := s.db.QueryRowContext(ctx, q, id).Scan(&name, &abv); err != nil {
		return beershop.Beer{}, err
	}

	return beershop.Beer{
		ID:   id,
		Name: name,
		Abv:  abv,
	}, nil
}
