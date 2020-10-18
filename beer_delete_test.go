package beershop_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func Test_Delete(t *testing.T) {
	ss := []struct {
		name    string
		cmd     beershop.DeleteBeerCmd
		errKeys []string
	}{
		{
			name:    "Invalid Id",
			cmd:     beershop.DeleteBeerCmd{ID: uuid.Nil},
			errKeys: []string{"ID"},
		},
		{
			name:    "Valid",
			cmd:     beershop.DeleteBeerCmd{ID: uuid.New()},
			errKeys: nil,
		},
	}

	for _, s := range ss {
		t.Run(s.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockRepository(ctrl)
			ct := 0
			if s.errKeys == nil {
				ct = 1
			}
			repo.EXPECT().Delete(ctx, s.cmd.ID).Return(nil).Times(ct)

			h := beershop.NewDeleteBeerHandler(repo)
			r, err := h(ctx, s.cmd)
			if err != nil && err != beershop.ErrValidationFailed {
				t.Fatalf("error running validation: %v", err)
			}

			for _, ek := range s.errKeys {
				vr := r.Validation
				if vr == nil {
					t.Fatalf("expected errors %v, none returned", vr)
				}

				if _, ok := vr.Errors()[ek]; !ok {
					t.Fatalf(`expected err "%s" not returned`, ek)
				}
			}
		})
	}
}
