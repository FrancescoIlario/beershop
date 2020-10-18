package grpc

import (
	context "context"
	"fmt"

	"github.com/FrancescoIlario/beershop"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Create(ctx context.Context, r *CreateBeerRequest) (*CreateBeerReply, error) {
	handleErr := func(cr *beershop.CreateBeerCmdResult, err error) error {
		switch err {
		case beershop.ErrValidationFailed:
			e := status.New(codes.InvalidArgument, fmt.Sprintf("Error handling request: %v", err))
			ve := ValidationError{
				ErrorMessage: err.Error(),
				Errors:       protoValidationError(cr.Validation),
			}
			e.WithDetails(&ve)
			return e.Err()
		case beershop.ErrConflict:
			return status.Error(codes.AlreadyExists, err.Error())
		default:
			return err
		}
	}

	cmd := beershop.CreateBeerCmd{Name: r.Name, Abv: r.Abv}
	cr, err := s.Be.Create(ctx, cmd)
	if err != nil {
		return nil, handleErr(cr, err)
	}

	return &CreateBeerReply{Id: cr.Result.ID.String()}, nil
}
