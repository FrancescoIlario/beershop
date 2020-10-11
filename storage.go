package beershop

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(context.Context, Beer) (uuid.UUID, error)
	Delete(context.Context, uuid.UUID) error
	List(context.Context) ([]Beer, error)
	Read(context.Context, uuid.UUID) (Beer, error)
}

var ErrNotFound = fmt.Errorf("beer not found")
