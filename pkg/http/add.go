package http

import (
	"btc/pkg/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	Amount   *float32 `json:"amount" form:"amount"`
	Datetime string   `json:"datetime" form:"amount"`
}

type AddService interface {
	Save(add *service.Add) error
}

type AddHandler struct {
	service AddService
}

func (a *AddHandler) Save(c echo.Context) error {
	add := new(AddRequest)

	if err := c.Bind(add); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if add.Amount == nil {
		return c.String(http.StatusBadRequest, "amount required")
	}

	if err := a.service.Save(&service.Add{Amount: *add.Amount, Datetime: add.Datetime}); err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.NoContent(http.StatusCreated)
}

func NewAddHandler(service AddService) *AddHandler {
	return &AddHandler{service}
}
