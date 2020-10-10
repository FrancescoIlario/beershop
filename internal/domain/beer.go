package domain

import "github.com/google/uuid"

// Beer ...
type Beer struct {
	Id   uuid.UUID
	Name string
	Abv  float32
}
