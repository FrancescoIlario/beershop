package rest

import (
	"log"
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) handleRead() http.HandlerFunc {
	type beer struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Abv  float32   `json:"abv"`
	}
	type response struct {
		Beer beer `json:"beer"`
	}
	handleErr := func(w http.ResponseWriter, r *http.Request, err error, cr *beershop.ReadBeerQryResult) {
		switch err {
		case beershop.ErrNotFound:
			s.respond(w, r, err.Error(), http.StatusNotFound)
		case beershop.ErrValidationFailed:
			s.invalid(w, r, cr.Validation.Errors())
		default:
			s.error(w, r, http.StatusInternalServerError, ErrCodeInternal, "Error handling request")
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
			s.respond(w, r, E{Message: "not valid id"}, http.StatusBadRequest)
			s.L.Logf("not valid id: %v", err)
			return
		}

		// reading from database
		qry := beershop.ReadBeerQry{ID: id}
		res, err := s.Be.Read(r.Context(), qry)
		if err != nil {
			handleErr(w, r, err, res)
			log.Printf("error reading data: %v", err)
			return
		}

		// responding
		s.respond(w, r, res, http.StatusOK)
	}
}
