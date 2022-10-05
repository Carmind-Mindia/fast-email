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
	inputs  <-chan amqp.Delivery
	manager manager.EmailManager
}

func NewRabbitMqEmail(channel *amqp.Channel) RabbitMqEmail {

	q, err := channel.QueueDeclare("notifications", false, true, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind(q.Name, "notification.*.ready", "carmind", false, nil)
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

	instance := RabbitMqEmail{inputs: messages}
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
		case "notification..ready":

			break
		}
		// For example, show received message in a console.
		log.Printf(" > Received message: %s\n", message.Body)
	}
}
