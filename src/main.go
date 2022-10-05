package main

import (
	"fmt"

	"github.com/Fonzeca/FastEmail/src/manager"
	"github.com/Fonzeca/FastEmail/src/service"
	"github.com/labstack/echo"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	e := echo.New()

	//Corremos el deamon con el channel
	go manager.DeamonEmail()
	go manager.DeamonNotification()
	channel, closeFunc := setupRabbitMq()
	defer closeFunc()

	//Creamos la api
	emailApi := service.NewApiEmail()
	service.NewRabbitMqEmail(channel)
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

func setupRabbitMq() (*amqp.Channel, func()) {
	// Create a new RabbitMQ connection.

	connectRabbitMQ, err := amqp.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		panic(err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		connectRabbitMQ.Close()
		panic(err)
	}

	return channelRabbitMQ, func() { connectRabbitMQ.Close(); channelRabbitMQ.Close() }
}

func InitConfig() {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
