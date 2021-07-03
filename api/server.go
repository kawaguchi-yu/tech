package main

import (
	"hello/server/infra"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 出力の定義

func main() {
	infra.DBInit()
	e := echo.New()
	e.Use(middleware.CORS())
	infra.Routing(e)
	e.Logger.Fatal(e.Start(":8080"))
}
