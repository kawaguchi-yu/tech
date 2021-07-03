package main

import (
	"hello/server/infra"

	"github.com/labstack/echo/v4"
)

// 出力の定義

func main() {
	infra.DBInit()
	e := echo.New()
	infra.Routing(e)
	e.Logger.Fatal(e.Start(":8080"))
}
