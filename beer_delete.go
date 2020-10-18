package beershop

import (
	"context"

	"github.com/google/uuid"
)

// DeleteBeerCmd command to delete a new beer
type DeleteBeerCmd struct {
	ID uuid.UUID
}

func (cmd *DeleteBeerCmd) Validate(ctx context.Context) (ValidationResult, error) {
	errs := make(map[string]string)

	if cmd.ID == uuid.Nil {
		errs["ID"] = "ID is not a valid UUID"
	}

	if len(errs) > 0 {
		return &validationResult{errors: errs}, ErrValidationFailed
	}
	return nil, nil
}

// DeleteBeerCmdResult the result of the delete operation
type DeleteBeerCmdResult struct {
	Validation ValidationResult
	Result     *struct {
		ID uuid.UUID
	}
}

// DeleteBeerHandlerFunc Handler for DeleteBeerCmd
type DeleteBeerHandlerFunc func(context.Context, DeleteBeerCmd) (*DeleteBeerCmdResult, error)

// NewDeleteBeerHandler DeleteBeerHandler constructor
func NewDeleteBeerHandler(repo Repository) DeleteBeerHandlerFunc {
	return func(ctx context.Context, cmd DeleteBeerCmd) (*DeleteBeerCmdResult, error) {
		if r, err := cmd.Validate(ctx); err != nil {
			return &DeleteBeerCmdResult{Validation: r}, err
		}

		if err := repo.Delete(ctx, cmd.ID); err != nil {
			return nil, err
		}
		return &DeleteBeerCmdResult{Result: &struct{ ID uuid.UUID }{ID: cmd.ID}}, nil
	}
}
