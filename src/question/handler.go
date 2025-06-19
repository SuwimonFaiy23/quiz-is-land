package question

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetQuestion(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetQuestion(c echo.Context) error {
	sessionID := c.Param("session_id")
	resp, err := h.service.GetQuestion(sessionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "get question failed"})
	}
	if resp.ID == 0 {
		if err := h.service.EndSession(sessionID); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "end session failed"})
		}
		return c.JSON(http.StatusOK, echo.Map{"question": nil})
	}
	return c.JSON(http.StatusOK, resp)
}


func (h *handler) SubmitAnswer(c echo.Context) error {
	var ans Answer
	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid answer"})
	}
	if err := h.service.SubmitAnswer(ans); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "cannot save answer"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}

func (h *handler) GetSummary(c echo.Context) error {
	sessionID := c.Param("sessionID")
	resp, err := h.service.GetSummary(sessionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "summary failed"})
	}
	return c.JSON(http.StatusOK, resp)
}