package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/http/rest"
	"github.com/FrancescoIlario/beershop/internal/storage/inmem"
)

func Test_Delete(t *testing.T) {
	// arrange
	st := inmem.New()
	id, err := st.Create(beershop.Beer{
		Name: "Beer 1",
		Abv:  1.0,
	})
	if err != nil {
		t.Fatalf("error arranging inmem db: %v", err)
	}
	sv := rest.NewServer(st)

	req := httptest.NewRequest(http.MethodDelete, "/beer/"+id.String(), nil)
	w := httptest.NewRecorder()

	// act
	sv.ServeHTTP(w, req)

	// assert
	if s := w.Result().StatusCode; s != http.StatusOK {
		t.Fatalf("expected %v, obtained %v", http.StatusOK, s)
	}
}
