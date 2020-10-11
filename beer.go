package beershop

import "github.com/google/uuid"

// Beer A Beer is a beer in the shop
type Beer struct {
	ID   uuid.UUID
	Name string
	Abv  float32
}
