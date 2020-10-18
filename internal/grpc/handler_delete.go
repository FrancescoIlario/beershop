package grpc

import (
	context "context"
	"fmt"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Delete(ctx context.Context, r *DeleteBeerRequest) (*DeleteBeerReply, error) {
	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Invalid uuid provided: %v", err))
	}

	cmd := beershop.DeleteBeerCmd{ID: id}
	rep, err := s.Be.Delete(ctx, cmd)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error handling request: %v", err))
	}

	return &DeleteBeerReply{
		Id: rep.Result.ID.String(),
	}, nil
}
