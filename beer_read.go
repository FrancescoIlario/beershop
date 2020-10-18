package beershop

import (
	"context"

	"github.com/google/uuid"
)

// ReadBeerQry command to read a new beer
type ReadBeerQry struct {
	ID uuid.UUID
}

func (qry *ReadBeerQry) Validate(ctx context.Context) (ValidationResult, error) {
	errs := make(map[string]string)

	if qry.ID == uuid.Nil {
		errs["ID"] = "ID is not a valid UUID"
	}

	if len(errs) > 0 {
		return &validationResult{errors: errs}, ErrValidationFailed
	}
	return nil, nil
}

// ReadBeerQryResult the result of the read operation
type ReadBeerQryResult struct {
	Validation ValidationResult
	Result     *struct {
		Beer ReadBeerQryBeerViewModel
	}
}

type ReadBeerQryBeerViewModel struct {
	ID   uuid.UUID
	Name string
	Abv  float32
}

// ReadBeerHandlerFunc Handler for ReadBeerQry
type ReadBeerHandlerFunc func(context.Context, ReadBeerQry) (*ReadBeerQryResult, error)

// NewReadBeerHandler ReadBeerHandler constructor
func NewReadBeerHandler(repo Repository) ReadBeerHandlerFunc {
	return func(ctx context.Context, qry ReadBeerQry) (*ReadBeerQryResult, error) {
		if r, err := qry.Validate(ctx); err != nil {
			return &ReadBeerQryResult{Validation: r}, err
		}

		b, err := repo.Read(ctx, qry.ID)
		if err != nil {
			return nil, err
		}
		br := ReadBeerQryBeerViewModel{
			ID:   b.ID,
			Abv:  b.Abv,
			Name: b.Name,
		}
		return &ReadBeerQryResult{Result: &struct{ Beer ReadBeerQryBeerViewModel }{Beer: br}}, nil
	}
}
