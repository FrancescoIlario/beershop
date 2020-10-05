package rest

import (
	"encoding/json"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error encoding response as JSON"))
		s.log.Logf("error encoding response as JSON: %v", err)
	}
}
