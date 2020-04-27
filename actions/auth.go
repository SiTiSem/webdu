package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var UserPassword *string

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	// Throws unauthorized error
	if u.Username != "admin" || u.Password != *UserPassword {
		return echo.ErrForbidden
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Server Tool"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func PasswordRandom() string {

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
