package user

import (
	"myauth/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func IndexHandler(c echo.Context) error {
	user, err := model.GetUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{user})
}

func StoreHandler(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
	}

	user.ID = bson.NewObjectId()

	if err = model.InsertUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	return c.JSON(http.StatusCreated, Response{user})
}
