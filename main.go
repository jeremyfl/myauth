package main

import (
	"myauth/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	r := e.Group("/user")
	// TODO:: replace "secret" to ENV
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", user.IndexHandler)
	r.POST("", user.StoreHandler)

	e.POST("/login", user.Login)

	e.Logger.Fatal(e.Start(":1323"))
}
