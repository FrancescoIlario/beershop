package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop/internal/storage"
	"github.com/FrancescoIlario/beershop/pkg/log"
)

type server struct {
	db  storage.Repository
	mux http.Handler
	log log.Logger
}

func NewServer(db storage.Repository) http.Handler {
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
