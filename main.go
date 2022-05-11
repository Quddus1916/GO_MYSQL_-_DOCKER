package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
    "github.com/Quddus1916/GO_MYSQL_DOCKER/controller"
)

func main(){
      e:= echo.New()
	  e.GET("/homepage",controller.homePage)
}