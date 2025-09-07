package app

import (
	"task-shooter-api/internal/app/server"
	"task-shooter-api/internal/app/server/handlers"

	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	e := echo.New()
	h := handlers.New()

	return server.New(e, h)
}
