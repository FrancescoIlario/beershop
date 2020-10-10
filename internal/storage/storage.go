package storage

import (
	"fmt"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(beershop.Beer) (uuid.UUID, error)
	Delete(uuid.UUID) error
	List() ([]beershop.Beer, error)
	Read(uuid.UUID) (beershop.Beer, error)
}

var ErrNotFound = fmt.Errorf("beer not found")
