package main

import (
	"api.com/db"
	"api.com/routers"
	"github.com/labstack/echo/v4"
)

func main() {

	db.InitDB()
	defer db.CloseDB()

	e := echo.New()

	routers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
