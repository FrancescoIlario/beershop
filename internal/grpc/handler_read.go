package grpc

import (
	context "context"
	"fmt"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Read(ctx context.Context, r *ReadBeerRequest) (*ReadBeerReply, error) {
	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Invalid uuid provided: %v", err))
	}

	qry := beershop.ReadBeerQry{ID: id}
	cr, err := s.Be.Read(ctx, qry)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error handling request: %v", err))
	}

	b := cr.Result.Beer
	return &ReadBeerReply{
		Beer: &Beer{
			Id:   b.ID.String(),
			Name: b.Name,
			Abv:  b.Abv,
		},
	}, nil
}
