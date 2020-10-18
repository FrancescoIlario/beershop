package rest

import (
	"fmt"
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) handleDelete() http.HandlerFunc {
	handleErr := func(w http.ResponseWriter, r *http.Request, err error, cr *beershop.DeleteBeerCmdResult) {
		switch err {
		case beershop.ErrNotFound:
			s.respond(w, r, err.Error(), http.StatusConflict)
		case beershop.ErrValidationFailed:
			s.invalid(w, r, cr.Validation.Errors())
		default:
			s.error(w, r, http.StatusInternalServerError, ErrCodeInternal, "Error handling request")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// decode request
		v := mux.Vars(r)
		idStr := v["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			m := fmt.Sprintf("provided id is not a valid guid: %v", idStr)
			s.respond(w, r, E{Message: m}, http.StatusBadRequest)
			s.L.Logf("provided id (%v) is not a valid guid: %v", idStr, err)
			return
		}

		// persisting request
		cmd := beershop.DeleteBeerCmd{ID: id}
		dr, err := s.Be.Delete(r.Context(), cmd)
		if err != nil {
			s.L.Logf("error handling request: %v", err)
			handleErr(w, r, err, dr)
			return
		}

		// responding
		s.respond(w, r, dr.Result, http.StatusOK)
	}
}
