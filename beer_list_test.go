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

func Test_List(t *testing.T) {
	qry := beershop.ListBeerQry{}
	is := is.New(t)
	ctx := context.Background()
	bb := []beershop.Beer{{ID: uuid.New(), Abv: 1.0, Name: "Name"}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().List(ctx).Return(bb, nil).Times(1)

	h := beershop.NewListBeerHandler(repo)
	r, err := h(ctx, qry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	bbr := []beershop.ListBeerQryBeerViewModel{
		{
			ID:   bb[0].ID,
			Abv:  bb[0].Abv,
			Name: bb[0].Name,
		},
	}
	is.Equal(r.Result.Beers, bbr)
}
