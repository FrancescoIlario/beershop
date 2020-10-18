package rest

import (
	"encoding/json"
	"net/http"
)

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error encoding response as JSON"))
		s.L.Logf("error encoding response as JSON: %v", err)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, status int, code ErrorCode, msg string) {
	ve := E{
		Code:    code,
		Message: msg,
	}
	s.respond(w, r, ve, status)
}

func (s *Server) invalid(w http.ResponseWriter, r *http.Request, validationErrors map[string]string) {
	ve := E{
		Code:       ErrCodeValidationFailed,
		Message:    "Validation Failed",
		Validation: validationErrors,
	}
	s.respond(w, r, ve, http.StatusBadRequest)
}
