package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setMiddlewares(e *echo.Echo) {
	// リクエストログを出力するようにする
	e.Use(middleware.Logger())

	// panicが発生してもサーバーを落とさずに500 Internal Server Errorを返すようにする
	e.Use(middleware.Recover())

	// このAPIを使えるフロントエンドを指定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://task-shooter.vercel.app", "http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
	}))
}
