package http

import (
	"btc/pkg/http/mock_http"
	service "btc/pkg/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {

	type query struct {
		Key   string
		Value string
	}
	type test struct {
		queries        []query
		listArguments  *service.ListServiceParams
		times          int
		err            error
		data           []*service.Wallet
		expectedStatus int
	}

	tests := []test{
		{queries: []query{{"startDatetime", "2022-12-29T20:43:07+0530"}}, times: 1, listArguments: &service.ListServiceParams{StartDatetime: "2022-12-29T20:43:07+0530", EndDatetime: ""}, err: nil, data: []*service.Wallet{}, expectedStatus: http.StatusOK},
		{queries: []query{{"startDatetime", "2022-12-29T20:43:07+0530"}, {"endDatetime", "2023-01-01T20:43:07+0530"}}, times: 1, listArguments: &service.ListServiceParams{StartDatetime: "2022-12-29T20:43:07+0530", EndDatetime: "2023-01-01T20:43:07+0530"}, err: nil, data: []*service.Wallet{}, expectedStatus: http.StatusOK},
		{queries: []query{{"startDatetime", "2022-12-29T20:43:07+0530"}}, times: 1, listArguments: &service.ListServiceParams{StartDatetime: "2022-12-29T20:43:07+0530", EndDatetime: ""}, err: errors.New("server error"), data: []*service.Wallet{}, expectedStatus: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		e := echo.New()
		q := make(url.Values)
		for _, qry := range tt.queries {
			q.Set(qry.Key, qry.Value)
		}
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctrl := gomock.NewController(t)
		s := mock_http.NewMockListService(ctrl)
		h := NewListHandler(s)

		s.EXPECT().List(gomock.Eq(tt.listArguments)).Times(tt.times).Return(tt.data, tt.err)

		// Assertions
		if assert.NoError(t, h.List(c)) {
			assert.Equal(t, tt.expectedStatus, rec.Code)
		}
	}

}
