package infra

import (
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) error {
	e.POST("/registrantion", func(c echo.Context) error {
		DBCreateUser(c, GetDB())
		return nil
	}) //user.structのデータを貰って登録する
	return nil
}
