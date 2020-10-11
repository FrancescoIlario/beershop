package rest

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *server) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode request
		v := mux.Vars(r)
		idStr := v["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			m := fmt.Sprintf("provided id is not a valid guid: %v", idStr)
			s.respond(w, r, e{Message: m}, http.StatusBadRequest)
			s.log.Logf("provided id (%v) is not a valid guid: %v", idStr, err)
			return
		}

		// persisting request
		if err := s.db.Delete(r.Context(), id); err != nil {
			s.respond(w, r, e{Message: "error persisting data into database"}, http.StatusInternalServerError)
			s.log.Logf("error persisting data into database: %v", err)
			return
		}

		// responding
		s.respond(w, r, nil, http.StatusOK)
	}
}
