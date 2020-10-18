package beershop_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func Test_Create(t *testing.T) {
	ss := []struct {
		name    string
		cmd     beershop.CreateBeerCmd
		errKeys []string
	}{
		{
			name: "Invalid Abv",
			cmd: beershop.CreateBeerCmd{
				Name: "Invalid Abv",
				Abv:  -1,
			},
			errKeys: []string{"Abv"},
		},
		{
			name: "Invalid Name",
			cmd: beershop.CreateBeerCmd{
				Name: "",
				Abv:  1.0,
			},
			errKeys: []string{"Name"},
		},
		{
			name: "Invalid Everything",
			cmd: beershop.CreateBeerCmd{
				Name: "",
				Abv:  -1,
			},
			errKeys: []string{"Name", "Abv"},
		},
		{
			name: "Valid",
			cmd: beershop.CreateBeerCmd{
				Name: "Valid Beer",
				Abv:  1.0,
			},
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
			repo.EXPECT().Create(ctx, gomock.Any()).Return(uuid.New(), nil).Times(ct)

			h := beershop.NewCreateBeerHandler(repo)
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
