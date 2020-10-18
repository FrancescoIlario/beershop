package beershop

import (
	"context"

	"github.com/google/uuid"
)

// ListBeerQry command to list a new beer
type ListBeerQry struct{}

// ListBeerQryResult the result of the list operation
type ListBeerQryResult struct {
	Result *struct {
		Beers []ListBeerQryBeerViewModel
	}
}

// ListBeerQryBeerViewModel the beer information returned in from the list commmand invokation
type ListBeerQryBeerViewModel struct {
	ID   uuid.UUID
	Name string
	Abv  float32
}

// ListBeerHandlerFunc Handler for ListBeerQry
type ListBeerHandlerFunc func(context.Context, ListBeerQry) (*ListBeerQryResult, error)

// NewListBeerHandler ListBeerHandler constructor
func NewListBeerHandler(repo Repository) ListBeerHandlerFunc {
	return func(ctx context.Context, qry ListBeerQry) (*ListBeerQryResult, error) {
		bb, err := repo.List(ctx)
		if err != nil {
			return nil, err
		}
		vbb := make([]ListBeerQryBeerViewModel, len(bb))
		for i, b := range bb {
			vbb[i] = ListBeerQryBeerViewModel{
				ID:   b.ID,
				Abv:  b.Abv,
				Name: b.Name,
			}
		}
		return &ListBeerQryResult{Result: &struct{ Beers []ListBeerQryBeerViewModel }{Beers: vbb}}, nil
	}
}
