package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FrancescoIlario/beershop/internal/http/rest"
	"github.com/FrancescoIlario/beershop/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func Test_Delete(t *testing.T) {
	// arrange
	is := is.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := uuid.New()
	st := mocks.NewMockRepository(mockCtrl)
	st.EXPECT().Delete(gomock.Any(), id).Return(nil).Times(1)

	sv := rest.NewServer(st)

	req := httptest.NewRequest(http.MethodDelete, "/beer/"+id.String(), nil)
	w := httptest.NewRecorder()

	// act
	sv.ServeHTTP(w, req)

	// assert
	is.Equal(w.Result().StatusCode, http.StatusOK)
}
