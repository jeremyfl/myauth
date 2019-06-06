package main

import (
	"myauth/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/login", controller.LoginHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
