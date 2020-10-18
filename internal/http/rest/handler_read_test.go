package rest_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	buildRequest := func(t *testing.T, id uuid.UUID) *http.Request {
		return httptest.NewRequest(http.MethodGet, "/beer/"+id.String(), nil)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cc := []struct {
		name       string
		statusCode int
		id         uuid.UUID
		errCode    rest.ErrorCode
		err        error
		res        *beershop.ReadBeerQryResult
	}{
		{
			name:       "valid",
			statusCode: http.StatusOK,
			id:         uuid.New(),
			err:        nil,
			res: &beershop.ReadBeerQryResult{
				Result: &struct {
					Beer beershop.ReadBeerQryBeerViewModel
				}{
					Beer: beershop.ReadBeerQryBeerViewModel{
						ID:   uuid.New(),
						Name: "first",
						Abv:  1.0,
					},
				},
			},
		},
		{
			name:       "invalid",
			statusCode: http.StatusBadRequest,
			id:         uuid.Nil,
			err:        beershop.ErrValidationFailed,
			errCode:    rest.ErrCodeValidationFailed,
			res: &beershop.ReadBeerQryResult{
				Validation: func() beershop.ValidationResult {
					vr := mocks.NewMockValidationResult(ctrl)
					vr.EXPECT().Errors().Return(map[string]string{
						"ID": "invalid ID",
					})
					return vr
				}(),
			},
		},
		{
			name:       "internal error",
			id:         uuid.New(),
			statusCode: http.StatusInternalServerError,
			errCode:    rest.ErrCodeInternal,
			err:        fmt.Errorf("internal error"),
			res:        nil,
		},
	}

	// arrange
	is := is.New(t)

	for _, c := range cc {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			be := mocks.NewMockBackend(ctrl)

			sv, td := NewTestServer(t, be)
			defer td()

			req := buildRequest(t, c.id)
			w := httptest.NewRecorder()

			be.
				EXPECT().
				Read(gomock.Any(), gomock.Any()).
				Return(c.res, c.err).
				Times(1)

			// act
			sv.ServeHTTP(w, req)

			// assert
			is.Equal(w.Result().StatusCode, c.statusCode)
			if w.Result().StatusCode >= 400 {
				b, err := ioutil.ReadAll(w.Result().Body)
				if err != nil {
					t.Fatal("error reading response body")
				}

				var e rest.E
				if err := json.Unmarshal(b, &e); err != nil {
					t.Fatal("error unmarshaling response body")
				}

				is.Equal(e.Code, c.errCode)
			}
		})
	}

}
