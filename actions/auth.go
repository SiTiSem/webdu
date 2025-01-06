package actions

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	length := 12
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}
