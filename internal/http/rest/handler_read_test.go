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

func Test_Read(t *testing.T) {
	// arrange
	is := is.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	st := mocks.NewMockRepository(mockCtrl)
	b := beershop.Beer{
		ID:   uuid.New(),
		Name: "Beer 1",
		Abv:  1.0,
	}
	st.EXPECT().Read(b.ID).Return(b, nil).Times(1)

	sv := rest.NewServer(st)
	req := httptest.NewRequest(http.MethodGet, "/beer/"+b.ID.String(), nil)
	w := httptest.NewRecorder()

	// act
	sv.ServeHTTP(w, req)

	// assert
	resp := struct {
		Beer struct {
			ID   uuid.UUID `json:"id"`
			Name string    `json:"name"`
			Abv  float32   `json:"abv"`
		} `json:"beer"`
	}{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("error reading JSON response: %v", err)
	}

	is.Equal(resp.Beer.ID, b.ID)
	is.Equal(resp.Beer.Name, b.Name)
	is.Equal(resp.Beer.Abv, b.Abv)
}
