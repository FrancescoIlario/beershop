package rest

import (
	"net/http"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/pkg/log"
)

//Server the beershop rest server
type Server struct {
	mux http.Handler

	L  log.Logger
	Be beershop.Backend
}

// NewServer the beershp rest server constructor
func NewServer(db beershop.Repository) *Server {
	s := &Server{
		L:  &log.L{},
		Be: beershop.NewBackend(db),
	}
	s.RegisterRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
