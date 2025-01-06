package actions

import (
	"embed"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"log"
	"net/http"
)

var FrontendFs embed.FS

func Route(e *echo.Echo) *echo.Echo {

	frontendStripped, err := fs.Sub(FrontendFs, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	e.Use(middleware.Gzip())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	api := e.Group("/api")
	api.POST("/login", Login)

	//api.Use(middleware.JWT([]byte("secret")))
	api.Use(echojwt.JWT([]byte("secret")))

	tools := api.Group("/tools")
	tools.GET("/scan", Scan)
	tools.POST("/get", Get)
	tools.GET("/dirtree", DirTree)

	e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(frontendStripped))))

	return e
}
