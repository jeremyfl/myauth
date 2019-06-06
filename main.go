package main

import (
	"myauth/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/user", controller.GetAllUser)
	e.POST("/user", controller.RegisterHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
