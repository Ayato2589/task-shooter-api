package server

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Health(c echo.Context) error
}

func setRoutes(e *echo.Echo, h Handlers) {
	// サーバーの稼働確認用のエンドポイント
	e.GET("/healthz", h.Health)
}
