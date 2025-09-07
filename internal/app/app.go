package app

import (
	"github.com/labstack/echo/v4"
)

type App interface {
	Start()
}

type app struct {
	server *echo.Echo
}

func New() App {
	return &app{server: InitServer()}
}

func (a *app) Start() {
	a.server.Logger.Fatal(a.server.Start(":8080"))
}
