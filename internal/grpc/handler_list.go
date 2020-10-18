package grpc

import (
	context "context"
	"fmt"

	"github.com/FrancescoIlario/beershop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) List(ctx context.Context, r *ListBeersRequest) (*ListBeersReply, error) {
	convert := func(bb []beershop.ListBeerQryBeerViewModel) []*Beer {
		cc := make([]*Beer, len(bb))
		for i, b := range bb {
			cc[i] = &Beer{
				Id:   b.ID.String(),
				Name: b.Name,
				Abv:  b.Abv,
			}
		}
		return cc
	}

	qry := beershop.ListBeerQry{}
	cr, err := s.Be.List(ctx, qry)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error handling request: %v", err))
	}

	return &ListBeersReply{Beers: convert(cr.Result.Beers)}, nil
}
