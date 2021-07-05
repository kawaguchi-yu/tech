package infra

import (
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) error {
	e.POST("/registrantion", CreateUser) //user.structのデータを貰って登録する
	return nil
}
