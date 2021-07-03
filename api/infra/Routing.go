package infra

import (
	"hello/server/controller"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) error {
	e.POST("/registrantion", controller.CreateUser)
	return nil
}
