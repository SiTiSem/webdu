package actions

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Route(e *echo.Echo) *echo.Echo {

	e.Use(middleware.Gzip())

	e.GET("/", echo.WrapHandler(http.FileServer(packr.New("index", "../frontend/dist/"))))
	e.GET("/js/*", echo.WrapHandler(http.StripPrefix("/js/", http.FileServer(packr.New("js", "../frontend/dist/js")))))
	e.GET("/css/*", echo.WrapHandler(http.StripPrefix("/css/", http.FileServer(packr.New("css", "../frontend/dist/css")))))

	api := e.Group("/api")
	api.POST("/login", Login)

	api.Use(middleware.JWT([]byte("secret")))

	tools := api.Group("/tools")
	tools.GET("/scan", Scan)
	tools.POST("/get", Get)
	tools.GET("/dirtree", DirTree)

	return e
}
