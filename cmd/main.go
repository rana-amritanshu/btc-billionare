package main

import (
	"btc/internal"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	internal.Run(e)
	e.Logger.Fatal(e.Start(":1323"))
}
