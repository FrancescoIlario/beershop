package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FrancescoIlario/beershop/internal/http/rest"
	"github.com/FrancescoIlario/beershop/internal/storage/inmem"
	"github.com/google/uuid"
)

func Test_Create(t *testing.T) {
	// arrange
	st := inmem.New()
	sv := rest.NewServer(st)

	r := struct {
		Name string  `json:"name"`
		Abv  float32 `json:"abv"`
	}{"name", 1.0}

	b, err := json.Marshal(&r)
	if err != nil {
		t.Fatalf("error generating JSON from request: %v", err)
	}
	br := bytes.NewReader(b)

	req := httptest.NewRequest(http.MethodPost, "/beer", br)
	w := httptest.NewRecorder()

	// act
	sv.ServeHTTP(w, req)

	// assert
	resp := struct {
		ID uuid.UUID `json:"id"`
	}{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("error reading JSON response: %v", err)
	}

	if resp.ID.Variant() == uuid.Invalid {
		t.Errorf("Invalid UUID returned: %s", resp.ID)
	}
}
