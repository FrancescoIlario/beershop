package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

func (s *Server) handleList() http.HandlerFunc {
	type beer struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Abv  float32   `json:"abv"`
	}
	type response struct {
		Beers []beer `json:"beers"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := s.Be.List(r.Context(), beershop.ListBeerQry{})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, ErrCodeInternal, "Error handling request")
			s.L.Logf("error reading data: %v", err)
			return
		}

		// responding
		s.respond(w, r, resp, http.StatusOK)
	}
}
