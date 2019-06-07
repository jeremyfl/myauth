package user

import (
	"myauth/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login handler
func Login(c echo.Context) (err error) {
	var email, password string = c.FormValue("email"), c.FormValue("password")

	user, err := model.FindUser(email, password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["level"] = user.Level
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
