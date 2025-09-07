package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://task-shooter.vercel.app", "http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello task-shooter")
	})

	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
