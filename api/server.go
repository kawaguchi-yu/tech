package main

import (
	"hello/server/infra"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 出力の定義

func main() {
	infra.DBInit()
	e := echo.New()                  //eにecho.New(echoを使っている)
	e.Use(middleware.CORS())         //echoはCORSを使う
	infra.Routing(e)                 //infraのRoutingを実行
	e.Logger.Fatal(e.Start(":8080")) //ポート解放している
}
