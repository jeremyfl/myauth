package main

import (
	"myauth/user"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/user", user.IndexHandler)
	e.POST("/user", user.StoreHandler)

	e.POST("/login", user.Login)

	e.Logger.Fatal(e.Start(":1323"))
}
