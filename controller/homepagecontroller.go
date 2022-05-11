package controller

import(
	"github.com/labstack/echo"
	"net/http"

)

func homePage ( c echo.Context) error{
     a:= "hello"
	return c.JSON(http.StatusOK, a)
}