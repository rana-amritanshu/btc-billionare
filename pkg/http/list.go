package http

import (
	"btc/pkg/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ListRequest struct {
	StartDatetime string `query:"startDatetime"`
	EndDatetime   string `query:"endDatetime"`
}

type Wallet struct {
	Amount   float64 `json:"amount"`
	Datetime string  `json:"datetime"`
}

type ListService interface {
	List(params *service.ListServiceParams) ([]*service.Wallet, error)
}

type ListHandler struct {
	service ListService
}

func (a *ListHandler) List(c echo.Context) error {
	request := new(ListRequest)

	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	var results []*Wallet

	wallets, err := a.service.List(&service.ListServiceParams{
		StartDatetime: request.StartDatetime,
		EndDatetime:   request.EndDatetime,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	for _, wallet := range wallets {
		results = append(results, &Wallet{
			Amount:   wallet.Amount,
			Datetime: wallet.Datetime,
		})
	}

	return c.JSON(http.StatusOK, wallets)
}

func NewListHandler(service ListService) *ListHandler {
	return &ListHandler{service}
}
