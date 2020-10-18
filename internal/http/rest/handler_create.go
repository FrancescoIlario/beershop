package rest

import (
	"encoding/json"
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

func (s *Server) handleCreate() http.HandlerFunc {
	type request struct {
		Name string  `json:"name"`
		Abv  float32 `json:"abv"`
	}
	type response struct {
		ID uuid.UUID `json:"id"`
	}

	handleErr := func(w http.ResponseWriter, r *http.Request, err error, cr *beershop.CreateBeerCmdResult) {
		switch err {
		case beershop.ErrConflict:
			s.error(w, r, http.StatusConflict, ErrCodeConflict, err.Error())
		case beershop.ErrValidationFailed:
			s.invalid(w, r, cr.Validation.Errors())
		default:
			s.error(w, r, http.StatusInternalServerError, ErrCodeInternal, "Error handling request")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// decode request
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, ErrCodeInternal, "error decoding request")
			s.L.Logf("error decoding body: %v", err)
			return
		}

		// persisting request
		cmd := beershop.CreateBeerCmd{Name: req.Name, Abv: req.Abv}
		cr, err := s.Be.Create(r.Context(), cmd)
		if err != nil {
			s.L.Logf("error handling request: %v", err)
			handleErr(w, r, err, cr)
			return
		}

		// responding
		s.respond(w, r, response{ID: cr.Result.ID}, http.StatusCreated)
	}
}
