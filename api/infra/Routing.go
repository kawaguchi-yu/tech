package infra

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	e.GET("/icon", func(c echo.Context) error {
		ReadCookieReturnIcon(c, GetDB())
		fmt.Printf("ReadCookie.return.icon\n")
		return nil
	})
	e.GET("/user", func(c echo.Context) error {
		ReadCookieReturnUser(c, GetDB())
		fmt.Printf("ReadCookie.return.user\n")
		return nil
	})
	e.GET("/getalluserpost", func(c echo.Context) error {
		ReturnAllUserPost(c, GetDB())
		fmt.Printf("return.getalluserpost\n")
		return nil
	})
	e.POST("/getuserpost", func(c echo.Context) error {
		ReadURLReturnUserPost(c, GetDB())
		fmt.Printf("return.getuserpost\n")
		return nil
	})
	e.POST("/post", func(c echo.Context) error {
		CreatePost(c, GetDB())
		fmt.Printf("post処理が呼ばれました\n")
		return nil
	})
	e.POST("/updatepost", func(c echo.Context) error {
		UpDataPost(c, GetDB())
		fmt.Printf("updatepost処理が呼ばれました\n")
		return nil
	})
	e.POST("/deletepost", func(c echo.Context) error {
		DeletePost(c, GetDB())
		fmt.Printf("deletepost処理が呼ばれました\n")
		return nil
	})
	e.POST("/signup", func(c echo.Context) error {
		DBCreateUser(c, GetDB())
		return nil
	}) //user.structのデータを貰って登録する
	e.POST("/login", func(c echo.Context) error { //ログイン
		Login(c, GetDB())
		return nil
	})
	e.GET("/logout", func(c echo.Context) error { //cookie
		Logout(c, GetDB())
		fmt.Printf("logout処理が呼ばれました\n")
		return nil
	})
	e.POST("/seticon", func(c echo.Context) error {
		SetIcon(c, GetDB())
		fmt.Printf("SetIcon処理が呼ばれました\n")
		return nil
	})
}
