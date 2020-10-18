package grpc_test

import (
	context "context"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/grpc"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cc := []struct {
		name       string
		r          *grpc.CreateBeerRequest
		cmdRes     *beershop.CreateBeerCmdResult
		beErr      error
		statusCode codes.Code
	}{
		{
			name: "valid",
			r:    &grpc.CreateBeerRequest{Abv: 1.0, Name: "beer"},
			cmdRes: &beershop.CreateBeerCmdResult{
				Result: &struct{ ID uuid.UUID }{ID: uuid.New()},
			},
			beErr: nil,
		},
		{
			name: "invalid",
			r:    &grpc.CreateBeerRequest{Abv: -1.0, Name: "beer"},
			cmdRes: &beershop.CreateBeerCmdResult{
				Validation: func() beershop.ValidationResult {
					vr := mocks.NewMockValidationResult(ctrl)
					vr.EXPECT().Errors().Return(map[string]string{
						"Abv": "Abv must be a positive number"})
					return vr
				}(),
			},
			beErr:      beershop.ErrValidationFailed,
			statusCode: codes.InvalidArgument,
		},
		{
			name:       "conflict",
			r:          &grpc.CreateBeerRequest{Abv: 1.0, Name: "beer"},
			cmdRes:     nil,
			beErr:      beershop.ErrConflict,
			statusCode: codes.AlreadyExists,
		},
	}

	for _, c := range cc {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			is := is.New(t)
			be := mocks.NewMockBackend(ctrl)
			s := grpc.Server{Be: be}
			ctx := context.TODO()

			be.
				EXPECT().
				Create(ctx, gomock.Any()).
				Return(c.cmdRes, c.beErr)

			// act
			l, err := s.Create(ctx, c.r)

			// assert
			if err != nil {
				s, ok := status.FromError(err)
				is.True(ok)
				is.True(s.Code() == c.statusCode)
				if c.cmdRes != nil && c.cmdRes.Validation != nil {
					is.True(s.Details() != nil)
				}
				return
			}

			is.True(l.Id == c.cmdRes.Result.ID.String())
		})
	}
}
