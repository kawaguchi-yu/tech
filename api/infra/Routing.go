package infra

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	// e.GET("/Login", func(c echo.Context) error {
	// 	if err := CheckToken(c); err != nil {
	// 		return err
	// 	}
	// 	return c.String(http.StatusOK, "OK! you're logined!")
	// })
	e.GET("/cookie/set/:name", func(c echo.Context) error {
		WriteCookie(c)
		fmt.Printf("WriteCookie")
		return nil
	})
	e.GET("/cookie/get", func(c echo.Context) error {
		ReadCookie(c)
		fmt.Printf("ReadCookie")
		return nil
	})
	e.POST("/registrantion", func(c echo.Context) error {
		DBCreateUser(c, GetDB())
		return nil
	}) //user.structのデータを貰って登録する
	e.POST("/login", func(c echo.Context) error {
		Login(c, GetDB())
		return nil
	}) //emailがdbにあればパスワードを検証して、合っていれば
}
