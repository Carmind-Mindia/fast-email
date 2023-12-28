package main

import (
	"log"

	"github.com/Carmind-Mindia/fast-email/v2/src/manager"
	"github.com/Carmind-Mindia/fast-email/v2/src/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()

	//Corremos el deamon con el channel
	go manager.DeamonEmail()
	// go manager.DeamonNotification()

	//Creamos la api
	emailApi := service.NewApiEmail()
	notificationApi := service.NewApiNotification()

	//Routeamos
	e.POST("/sendRecoverPassword", emailApi.SendRecoverPassword)

	e.POST("/sendNotificationToCarmind", notificationApi.SendNotificationToCarmind)

	//Start!
	e.Logger.Fatal(e.Start(":5896"))
}
