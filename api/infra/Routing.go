package infra

import (
	"fmt"
	"hello/server/interfaces/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	userController := controllers.NewUserController(DBInit())
	e.POST("/updateuser", func(c echo.Context) error {
		fmt.Printf("/updateuser\n")
		return userController.UpdateUser(c)
	})
	e.POST("/seticon", func(c echo.Context) error {
		fmt.Printf("/seticon\n")
		return userController.SetIcon(c)
	})
	e.POST("/updatepost", func(c echo.Context) error { //試してない
		fmt.Printf("/updatepost\n")
		return userController.UpdatePost(c)
	})
	e.POST("/deletepost", func(c echo.Context) error {
		fmt.Printf("/deletepost\n")
		return userController.DeletePost(c)
	})
	e.POST("/post", func(c echo.Context) error {
		fmt.Printf("post処理が呼ばれました\n")
		return userController.CreatePost(c)
	})
	e.POST("/getuserpost", func(c echo.Context) error {
		fmt.Printf("/getuserpost\n")
		return userController.ReturnUserPostByName(c)
	})
	e.GET("/getalluserpost", func(c echo.Context) error {
		fmt.Printf("/getalluserpost\n")
		return userController.ReturnAllUserPost(c)
	})
	e.GET("/logout", func(c echo.Context) error { //cookie
		fmt.Printf("/logout\n")
		return userController.Logout(c)
	})
	e.GET("/user", func(c echo.Context) error {
		fmt.Printf("/user\n")
		return userController.ReadCookieReturnUser(c)
	})
	e.POST("/getuserbyid", func(c echo.Context) error {
		fmt.Printf("/getuserbyid\n")
		return userController.ReturnUserAndPostByPostID(c)
	})
	e.POST("/returngoodedpost", func(c echo.Context) error {
		fmt.Printf("/returngoodedpost\n")
		return userController.ReturnGoodedPost(c)
	})
	e.POST("/returngoodedpostbyword", func(c echo.Context) error {
		fmt.Printf("/returngoodedpostbyword\n")
		return userController.ReturnGoodedPostByWord(c)
	})
	e.GET("/guestlogin", func(c echo.Context) error {
		fmt.Printf("/guestlogin\n")
		return userController.GuestLogin(c)
	})
	e.POST("/login", func(c echo.Context) error {
		fmt.Printf("/login\n")
		return userController.Login(c)
	})
	e.POST("/good", func(c echo.Context) error {
		fmt.Printf("/good\n")
		return userController.CreateGood(c)
	})
	e.POST("/deletegood", func(c echo.Context) error {
		fmt.Printf("/deletegood\n")
		return userController.DeleteGoodByPostID(c)
	})
	e.POST("/signup", func(c echo.Context) error {
		fmt.Printf("/signup\n")
		return userController.CreateUser(c)
	}) //user.structのデータを貰って登録する
	e.GET("/deleteuser", func(c echo.Context) error {
		fmt.Printf("/deleteuser\n")
		return userController.DeleteUser(c)
	})
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}
