package controller

import (
	"fmt"
	"hello/server/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	u := new(domain.User)
	c.Bind(u)
	fmt.Printf("%+v\n\n", u)
	return c.JSON(http.StatusOK, "name:"+u.Name+", email:"+u.EMail+", password:"+u.Password)
}
