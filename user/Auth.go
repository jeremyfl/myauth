package user

import (
	"myauth/model"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) (err error) {
	var email, password string = c.FormValue("email"), c.FormValue("password")

	user, err := model.FindUser(email, password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
	}

	return c.JSON(http.StatusOK, Response{user})
}
