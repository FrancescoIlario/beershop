package beershop

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

// CreateBeerCmd command to create a new beer
type CreateBeerCmd struct {
	Name string
	Abv  float32
}

func (cmd *CreateBeerCmd) Validate(ctx context.Context) (ValidationResult, error) {
	errs := make(map[string]string)

	if cn := strings.TrimLeft(cmd.Name, " "); len(cn) == 0 {
		errs["Name"] = "Name is empty or composed only by whitespaces"
	}
	if cmd.Abv <= 0 {
		errs["Abv"] = "Abv must be a positive number"
	}

	if len(errs) > 0 {
		return &validationResult{errors: errs}, ErrValidationFailed
	}
	return nil, nil
}

// CreateBeerCmdResult the result of the create operation
type CreateBeerCmdResult struct {
	Validation ValidationResult
	Result     *struct {
		ID uuid.UUID
	}
}

// CreateBeerHandlerFunc Handler for CreateBeerCmd
type CreateBeerHandlerFunc func(context.Context, CreateBeerCmd) (*CreateBeerCmdResult, error)

// NewCreateBeerHandler NewCreateBeerHandler handler func constructor
func NewCreateBeerHandler(repo Repository) CreateBeerHandlerFunc {
	return func(ctx context.Context, cmd CreateBeerCmd) (*CreateBeerCmdResult, error) {
		if r, err := cmd.Validate(ctx); err != nil {
			return &CreateBeerCmdResult{Validation: r}, err
		}

		b := Beer{Name: cmd.Name, Abv: cmd.Abv}
		id, err := repo.Create(ctx, b)
		if err != nil {
			return nil, err
		}
		return &CreateBeerCmdResult{Result: &struct{ ID uuid.UUID }{ID: id}}, nil
	}
}
