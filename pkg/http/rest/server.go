package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop/pkg/storage"
)

type server struct {
	db  storage.Repository
	mux http.Handler
}

func NewServer(db storage.Repository) http.Handler {
	s := &server{
		db: db,
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
