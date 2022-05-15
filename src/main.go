package main

import (
	"fast-email/src/manager"
	"fast-email/src/service"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//Corremos el deamon con el channel
	go manager.Deamon()

	//Creamos la api
	emailApi := service.NewApiEmail()

	//Routeamos
	e.POST("/sendRecoverPassword", emailApi.SendRecoverPassword)

	//Start!
	e.Logger.Fatal(e.Start(":5896"))
}