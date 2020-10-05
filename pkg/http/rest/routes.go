package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) routes() {
	s.log.Logln("Registering routes")
	router := mux.NewRouter()
	{ // Beer
		router.HandleFunc("/beer", s.handleCreate()).Methods(http.MethodPost)
		router.HandleFunc("/beer", s.handleList()).Methods(http.MethodGet)
		router.HandleFunc("/beer/{id}", s.handleDelete()).Methods(http.MethodDelete)
		router.HandleFunc("/beer/{id}", s.handleRead()).Methods(http.MethodGet)
	}

	s.mux = router
}
