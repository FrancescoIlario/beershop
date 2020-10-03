package rest

import (
	"log"
	"net/http"

	"github.com/FrancescoIlario/beershop/pkg/domain"
	"github.com/google/uuid"
)

func (s *server) handleList() http.HandlerFunc {
	type beer struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Abv  float32   `json:"abv"`
	}
	type response struct {
		Beers []beer `json:"beers"`
	}
	convert := func(b []domain.Beer) []beer {
		lb := make([]beer, len(b))
		for i, be := range b {
			lb[i] = beer{
				Id:   be.Id,
				Name: be.Name,
				Abv:  be.Abv,
			}
		}
		return lb
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// reading from database
		b, err := s.db.List()
		if err != nil {
			s.respond(w, r, e{Message: "error reading data from database"}, http.StatusInternalServerError)
			log.Printf("error reading data from database: %v", err)
			return
		}

		// responding
		s.respond(w, r, response{Beers: convert(b)}, http.StatusOK)
	}
}
