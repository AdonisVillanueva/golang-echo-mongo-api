package main

import (
	"github.com/AdonisVillanueva/golang-echo-mongo-api/configs"
	"github.com/AdonisVillanueva/golang-echo-mongo-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
	})

	//run database
	configs.ConnectDB()
	//routes
	routes.UserRoute(e) //add this

	e.Logger.Fatal(e.Start(":6000"))
}
