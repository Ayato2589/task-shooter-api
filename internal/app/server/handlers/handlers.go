package handlers

import (
	"net/http"
	"task-shooter-api/internal/app/server"

	"github.com/labstack/echo/v4"
)

type handlers struct{}

func New() server.Handlers {
	return &handlers{}
}

func (h *handlers) Health(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
