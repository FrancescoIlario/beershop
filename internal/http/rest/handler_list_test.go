package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/http/rest"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func Test_List(t *testing.T) {
	// arrange
	is := is.New(t)

	bb := []beershop.Beer{
		{ID: uuid.New(), Name: "Beer 1", Abv: 1.0},
		{ID: uuid.New(), Name: "Beer 2", Abv: 2.0},
	}
	mbb := make(map[uuid.UUID]beershop.Beer)
	for _, b := range bb {
		mbb[b.ID] = b
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	st := mocks.NewMockRepository(mockCtrl)
	st.EXPECT().List(gomock.Any()).Return(bb, nil).Times(1)

	sv := rest.NewServer(st)
	req := httptest.NewRequest(http.MethodGet, "/beer", nil)
	w := httptest.NewRecorder()

	// act
	sv.ServeHTTP(w, req)

	// assert
	resp := struct {
		Beers []struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
			Abv  float32   `json:"abv"`
		} `json:"beers"`
	}{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("error reading JSON response: %v", err)
	}

	for _, b := range resp.Beers {
		mb, ok := mbb[b.ID]
		if !ok {
			t.Errorf("not expected beer id: %v", b.ID)
			continue
		}

		is.Equal(b.ID, mb.ID)
		is.Equal(b.Name, mb.Name)
		is.Equal(b.Abv, mb.Abv)
	}
	is.True(len(resp.Beers) == len(bb))
}
