package storage

import (
	"fmt"

	"github.com/FrancescoIlario/beershop/internal/domain"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(domain.Beer) (uuid.UUID, error)
	Delete(uuid.UUID) error
	List() ([]domain.Beer, error)
	Read(uuid.UUID) (domain.Beer, error)
}

var ErrNotFound = fmt.Errorf("beer not found")
