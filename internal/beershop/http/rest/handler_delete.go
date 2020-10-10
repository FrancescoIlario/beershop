package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *server) handleDelete() http.HandlerFunc {
	type request struct {
		Id uuid.UUID `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// decode request
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			s.respond(w, r, e{Message: "error deconding request"}, http.StatusBadRequest)
			s.log.Logf("error deconding body: %v", err)
			return
		}

		// persisting request
		if err := s.db.Delete(req.Id); err != nil {
			s.respond(w, r, e{Message: "error persisting data into database"}, http.StatusInternalServerError)
			s.log.Logf("error persisting data into database: %v", err)
			return
		}

		// responding
		s.respond(w, r, nil, http.StatusOK)
	}
}
