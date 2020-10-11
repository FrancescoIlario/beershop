package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

func (s *server) handleList() http.HandlerFunc {
	type beer struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Abv  float32   `json:"abv"`
	}
	type response struct {
		Beers []beer `json:"beers"`
	}
	convert := func(b []beershop.Beer) []beer {
		lb := make([]beer, len(b))
		for i, be := range b {
			lb[i] = beer{
				ID:   be.ID,
				Name: be.Name,
				Abv:  be.Abv,
			}
		}
		return lb
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// reading from database
		b, err := s.db.List(r.Context())
		if err != nil {
			s.respond(w, r, e{Message: "error reading data from database"}, http.StatusInternalServerError)
			s.log.Logf("error reading data from database: %v", err)
			return
		}

		// responding
		s.respond(w, r, response{Beers: convert(b)}, http.StatusOK)
	}
}
