package rest_test

import (
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/http/rest"
	"github.com/FrancescoIlario/beershop/pkg/log"
)

func NewTestServer(t *testing.T, be beershop.Backend) (*rest.Server, func()) {
	srv := rest.Server{
		Be: be,
		L:  &log.L{},
	}
	srv.RegisterRoutes()
	teardown := func() {}

	return &srv, teardown
}
