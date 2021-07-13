package infra

import (
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	// e.GET("/Login", func(c echo.Context) error {
	// 	if err := CheckToken(c); err != nil {
	// 		return err
	// 	}
	// 	return c.String(http.StatusOK, "OK! you're logined!")
	// })
	e.GET("/cookie/set/:name", WriteCookie)
	e.GET("/cookie/get", ReadCookie)
	e.POST("/registrantion", func(c echo.Context) error {
		DBCreateUser(c, GetDB())
		return nil
	}) //user.structのデータを貰って登録する
}
