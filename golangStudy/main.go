package main

import (
	"net/http"
	"strconv"

	"github.com/junstory/golangStudy/ex3"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/:number", func(c echo.Context) error {
		nstr := c.Param("number")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(ex3.IsPrime(n)))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
