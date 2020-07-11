package api

import "fmt"

type API struct {
	Router *echo.Echo
}

var (
	Apix API
)

func Init() *API {
	router := echo.New()

	return &Apix
}

func (a *API) SetupRoutes() {
	// Routes
	rpub := a.Router.Group("/public")
	pv1.GET("/timestamp", a.GetServerTime)
}
