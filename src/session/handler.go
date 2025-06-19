package session

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateSession(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateSession(c echo.Context) error {
	resp, err := h.service.CreateSession()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "cannot create session",
		})
	}

	return c.JSON(http.StatusOK, resp)
}
