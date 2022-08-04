package main

import (
	"github.com/Fonzeca/FastEmail/src/manager"
	"github.com/Fonzeca/FastEmail/src/service"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//Corremos el deamon con el channel
	go manager.DeamonEmail()

	//Creamos la api
	emailApi := service.NewApiEmail()
	notificationApi := service.NewApiNotification()

	//Routeamos
	e.POST("/sendRecoverPassword", emailApi.SendRecoverPassword)
	e.POST("/sendDocsCloseToExpire", emailApi.SendDocsCloseToExpire)
	e.POST("/sendNoneDocsCloseToExpire", emailApi.SendNoneDocsCloseToExpire)
	e.POST("/sendFailureEvaluacion", emailApi.SendFailureEvaluacion)

	e.POST("/sendNotificationToCarmind", notificationApi.SendNotificationToCarmind)

	//Start!
	e.Logger.Fatal(e.Start(":5896"))
}
