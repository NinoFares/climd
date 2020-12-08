package main

import (
	"db"
	"handler"
	"store"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//e.Validator = NewValidator()
	return e
}

func main() {
	r := New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	ps := store.NewPatientStore(d)
	h := handler.NewHandler(ps)

	h.Register(v1)
	r.Logger.Fatal(r.Start("0.0.0.0:3000"))
}
