package main

import(
	
	"github.com/labstack/echo"
    "github.com/Quddus1916/GO_MYSQL_DOCKER/controller"
	"github.com/labstack/echo/middleware"
)

func main(){
      e:= echo.New()
	  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " uri=${uri},method=${method}, status=${status}\n",
	  }))
	  e.GET("/homepage",controller.homePage)

	  e.Start(":8080")

}