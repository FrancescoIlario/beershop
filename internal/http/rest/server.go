package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/pkg/log"
)

type server struct {
	db  beershop.Repository
	mux http.Handler
	log log.Logger
}

func NewServer(db beershop.Repository) http.Handler {
	s := &server{
		db:  db,
		log: &log.L{},
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
