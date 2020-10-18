package beershop_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func Test_Read(t *testing.T) {
	ss := []struct {
		name    string
		qry     beershop.ReadBeerQry
		errKeys []string
	}{
		{
			name:    "Invalid Id",
			qry:     beershop.ReadBeerQry{ID: uuid.Nil},
			errKeys: []string{"ID"},
		},
		{
			name:    "Valid",
			qry:     beershop.ReadBeerQry{ID: uuid.New()},
			errKeys: nil,
		},
	}

	for _, s := range ss {
		t.Run(s.name, func(t *testing.T) {
			is := is.New(t)
			ctx := context.Background()
			b := beershop.Beer{ID: s.qry.ID, Abv: 1.0, Name: "Name"}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockRepository(ctrl)
			ct := 0
			if s.errKeys == nil {
				ct = 1
			}
			repo.EXPECT().Read(ctx, s.qry.ID).Return(b, nil).Times(ct)

			h := beershop.NewReadBeerHandler(repo)
			r, err := h(ctx, s.qry)
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

			if err == nil {
				rb := r.Result.Beer
				is.Equal(b.ID, rb.ID)
				is.Equal(b.Abv, rb.Abv)
				is.Equal(b.Name, rb.Name)
			}
		})
	}
}
