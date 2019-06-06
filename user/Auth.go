package user

import (
	"myauth/model"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) (err error) {
	u := new(model.User)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	user, err := model.FindUser(u)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
	}

	return c.JSON(http.StatusOK, Response{user})
}
