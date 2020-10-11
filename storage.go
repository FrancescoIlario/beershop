package beershop

import (
	"fmt"

	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(Beer) (uuid.UUID, error)
	Delete(uuid.UUID) error
	List() ([]Beer, error)
	Read(uuid.UUID) (Beer, error)
}

var ErrNotFound = fmt.Errorf("beer not found")
