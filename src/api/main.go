package main

import (
	"github.com/kynmh69/study-passkey/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server.Start(e)
}
