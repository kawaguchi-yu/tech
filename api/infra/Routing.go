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

	// e.GET("/set/:name", func(c echo.Context) error {
	// 	CreateCookie(c)
	// 	fmt.Printf("WriteCookie\n")
	// 	return nil
	// })
	e.GET("/user", func(c echo.Context) error {
		UserVerify(c, GetDB())
		fmt.Printf("ReadCookie\n")
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
	e.POST("/seticon", func(c echo.Context) error {
		SetIcon(c, GetDB())
		fmt.Printf("SetIcon\n")
		return nil
	})
}
