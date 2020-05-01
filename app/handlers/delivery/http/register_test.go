package http_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/napapol-dev-group/poc-golang/app/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_http "github.com/napapol-dev-group/poc-golang/app/handlers/delivery/http"
	_mockUseCase "github.com/napapol-dev-group/poc-golang/app/mocks/user"
)

func TestHandler_Register(t *testing.T) {
	requestURL := "/user/register"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUseCase := _mockUseCase.NewMockUseCase(ctrl)

	t.Run("Happy", func(t *testing.T) {

		jsonReqBody := `{
			"fullName": "napapol eiei",
			"address": "here"
		}`

		expectedResult := &entities.CreateUser{
			FullName: "napapol eiei",
			Address:  "here",
		}

		mockCreateUser := entities.CreateUser{
			FullName: "napapol eiei",
			Address:  "here",
		}

		mockUseCase.EXPECT().RegisterUser(mockCreateUser).Return(expectedResult, nil)

		gin.SetMode(gin.TestMode)
		ginEngine := gin.Default()
		_http.NewEndpointHTTPHandler(ginEngine, mockUseCase)

		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(jsonReqBody))
		assert.NoError(t, err)

		ginEngine.ServeHTTP(res, req)
		httpmock.Activate()

		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("Fail, use case return error", func(t *testing.T) {
		jsonReqBody := `{
			"fullName": "napapol eiei",
			"address": "here"
		}`

		mockCreateUser := entities.CreateUser{
			FullName: "napapol eiei",
			Address:  "here",
		}

		mockUseCase.EXPECT().RegisterUser(mockCreateUser).Return(nil, errors.New("some thing wrong"))

		gin.SetMode(gin.TestMode)
		ginEngine := gin.Default()
		_http.NewEndpointHTTPHandler(ginEngine, mockUseCase)

		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(jsonReqBody))
		assert.NoError(t, err)

		ginEngine.ServeHTTP(res, req)
		httpmock.Activate()

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

}
