package infra

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context) error {
	name := c.Param("name")
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = name
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie: "+name)
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie: "+cookie.Value)
}
