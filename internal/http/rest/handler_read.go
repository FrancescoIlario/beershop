package rest

import (
	"log"
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *server) handleRead() http.HandlerFunc {
	type beer struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Abv  float32   `json:"abv"`
	}
	type response struct {
		Beer beer `json:"beer"`
	}
	convert := func(b beershop.Beer) beer {
		return beer{
			Id:   b.Id,
			Name: b.Name,
			Abv:  b.Abv,
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}
		id, err := uuid.Parse(idStr)
		if err != nil {
			s.respond(w, r, e{Message: "not valid id"}, http.StatusBadRequest)
			s.log.Logf("not valid id: %v", err)
			return
		}

		// reading from database
		b, err := s.db.Read(id)
		if err != nil {
			s.respond(w, r, e{Message: err.Error()}, http.StatusInternalServerError)
			log.Printf("error reading data from database: %v", err)
			return
		}

		// responding
		s.respond(w, r, response{Beer: convert(b)}, http.StatusOK)
	}
}
