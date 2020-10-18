package rest_test

import (
	"bytes"
	"encoding/json"
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

func Test_CreateHandler(t *testing.T) {
	type request struct {
		Name string  `json:"name"`
		Abv  float32 `json:"abv"`
	}

	buildRequest := func(t *testing.T, body request) *http.Request {
		b, err := json.Marshal(&body)
		if err != nil {
			t.Fatalf("error generating JSON from request: %v", err)
		}
		br := bytes.NewReader(b)

		return httptest.NewRequest(http.MethodPost, "/beer", br)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cc := []struct {
		name       string
		req        request
		statusCode int
		errCode    rest.ErrorCode
		res        *beershop.CreateBeerCmdResult
		err        error
	}{
		{
			name:       "valid",
			req:        request{"name", 1.0},
			statusCode: http.StatusCreated,
			res: &beershop.CreateBeerCmdResult{
				Result: &struct {
					ID uuid.UUID
				}{ID: uuid.New()},
			},
			err: nil,
		},
		{
			name:       "invalid name",
			req:        request{"", 1.0},
			statusCode: http.StatusBadRequest,
			errCode:    rest.ErrCodeValidationFailed,
			res: &beershop.CreateBeerCmdResult{
				Validation: func() beershop.ValidationResult {
					vr := mocks.NewMockValidationResult(ctrl)
					errors := map[string]string{"Name": "invalid name"}
					vr.EXPECT().Errors().Return(errors).Times(1)
					return vr
				}(),
			},
			err: beershop.ErrValidationFailed,
		},
		{
			name:       "conflict",
			req:        request{"name", 1.0},
			statusCode: http.StatusConflict,
			errCode:    rest.ErrCodeConflict,
			err:        beershop.ErrConflict,
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

			req := buildRequest(t, c.req)
			w := httptest.NewRecorder()

			be.EXPECT().Create(gomock.Any(), gomock.Any()).Return(c.res, c.err).Times(1)

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
