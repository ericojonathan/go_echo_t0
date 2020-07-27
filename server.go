package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type API struct {
	Router *echo.Echo
}

var (
	Apix API
)

func main() {
	Init()
	Apix.Run()
}

func Init() /*API*/ {
	router := echo.New()
	Apix.Router = router
	//return &Apix
}

func (a API) Run() {
	a.Router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	a.Router.Logger.Fatal(a.Router.Start(":1323"))
}
