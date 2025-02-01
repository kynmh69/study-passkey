package server

import "github.com/labstack/echo/v4"

func Start(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}
