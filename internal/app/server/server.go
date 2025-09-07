package server

import "github.com/labstack/echo/v4"

func New(e *echo.Echo, h Handlers) *echo.Echo {
	setMiddlewares(e)
	setRoutes(e, h)

	return e
}
