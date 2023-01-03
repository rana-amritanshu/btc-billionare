package http

import (
	"btc/pkg/http/mock_http"
	service "btc/pkg/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {

	type test struct {
		request        string
		saveArguments  *service.Add
		times          int
		saveReturn     error
		expectedStatus int
	}

	tests := []test{
		{request: `{"amount": 10.10}`, times: 1, saveArguments: &service.Add{Amount: 10.1}, saveReturn: nil, expectedStatus: http.StatusCreated},
		{request: `{"amount": 10.10, "datetime": "2022-12-29T20:43:07+0530"}`, times: 1, saveArguments: &service.Add{Amount: 10.1, Datetime: "2022-12-29T20:43:07+0530"}, saveReturn: nil, expectedStatus: http.StatusCreated},
		{request: ``, saveArguments: &service.Add{}, times: 0, saveReturn: nil, expectedStatus: http.StatusBadRequest},
		{request: `{"amount": 10.10}`, times: 1, saveArguments: &service.Add{Amount: 10.1}, saveReturn: errors.New("Save Error"), expectedStatus: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.request))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctrl := gomock.NewController(t)
		s := mock_http.NewMockAddService(ctrl)
		h := NewAddHandler(s)

		s.EXPECT().Save(gomock.Eq(tt.saveArguments)).Times(tt.times).Return(tt.saveReturn)

		// Assertions
		if assert.NoError(t, h.Save(c)) {
			assert.Equal(t, tt.expectedStatus, rec.Code)
		}
	}

}
