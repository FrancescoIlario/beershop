package grpc

import (
	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/pkg/log"
)

type Server struct {
	Be beershop.Backend
	L  log.L
}
