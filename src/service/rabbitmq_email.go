package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Fonzeca/FastEmail/src/manager"
	"github.com/Fonzeca/FastEmail/src/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqEmail struct {
	inputs              <-chan amqp.Delivery
	manager             manager.EmailManager
	notificationManager manager.NotificationManager
}

func NewRabbitMqEmail(channel *amqp.Channel) RabbitMqEmail {

	q, err := channel.QueueDeclare("notifications", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(q.Name, "notification.#.ready", "carmind", false, nil)
	if err != nil {
		panic(err)
	}

	// Subscribing to QueueService1 for getting messages.
	messages, err := channel.Consume(
		q.Name,      // queue name
		"fastemail", // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // arguments
	)
	if err != nil {
		log.Println(err)
	}

	instance := RabbitMqEmail{
		inputs:              messages,
		notificationManager: manager.NewNotificationManager(),
	}
	go instance.Run()
	return instance
}

func (m *RabbitMqEmail) Run() {
	for message := range m.inputs {

		switch message.RoutingKey {
		case "notification.failure.ready":
			pojo := model.FailureEvaluacion{}
			err := json.Unmarshal(message.Body, &pojo)
			if err != nil {
				fmt.Println(err)
				break
			}
			m.manager.SendFailureEvaluacion(pojo)
			break
		case "notification.recover.password.ready":
			pojo := model.RecuperarContraseña{}
			err := json.Unmarshal(message.Body, &pojo)
			if err != nil {
				fmt.Println(err)
				break
			}
			m.manager.SendRecoverPassword(pojo)
			break
		case "notification.weekly.email.ready":
			pojo := model.ResumenSemanalLleno{}
			err := json.Unmarshal(message.Body, &pojo)
			if err != nil {
				fmt.Println(err)
				break
			}

			//Verificamos si es con vencimientos o no
			if pojo.Vencimientos != nil && len(pojo.Vencimientos) > 0 {
				m.manager.SendDocsCloseToExpire(pojo)
			} else {
				pojoVacio := model.ResumenSemanalVacio{
					Email:  pojo.Email,
					Nombre: pojo.Nombre,
				}
				m.manager.SendNoneDocsCloseToExpire(pojoVacio)
			}
			break
		case "notification.zone.fastemail.ready":
			pojo := model.ZoneNotification{}
			err := json.Unmarshal(message.Body, &pojo)
			if err != nil {
				fmt.Println(err)
				break
			}

			var message string = ""
			if pojo.EventType == "entra" {
				message = fmt.Sprintf("El vehículo '%s' entro de la zona '%s'", pojo.VehiculoName, pojo.ZoneName)
			} else if pojo.EventType == "sale" {
				message = fmt.Sprintf("El vehículo '%s' salio de la zona '%s'", pojo.VehiculoName, pojo.ZoneName)
			} else {
				message = "Error, event type desconocido"
			}

			notification := model.SimpleNotification{
				Title:   "Aviso de zona",
				Message: message,
				To:      pojo.FCMTokens,
			}

			m.notificationManager.SendNotificationToCarmind(notification)
			break
		}
		// For example, show received message in a console.
		log.Printf(" > Received message: %s\n", message.Body)
	}
}
