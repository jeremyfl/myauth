package controller

import (
	"myauth/model"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterHandler(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return nil
	}

	if err = model.InsertUser(user); err != nil {
		return nil
	}

	return c.JSON(http.StatusCreated, user)
}

func GetAllUser(c echo.Context) error {
	user, err := model.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
