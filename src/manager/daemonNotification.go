package manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Fonzeca/FastEmail/src/model"
	"gopkg.in/maddevsio/fcm.v1"
)

var (
	NotificationChannel chan model.CarmindNotification
)

var listMapNotification = make(map[string]time.Time)

func DeamonNotification() {
	//Creamos el channel
	NotificationChannel = make(chan model.CarmindNotification)

	for {

		//Esperamos un dato del canal
		data := <-NotificationChannel

		err := processNotification(data.To)
		if err != nil {
			//TODO: logeamos el error
			fmt.Print(err)
			continue
		}

		client := fcm.NewFCM(os.Getenv("FCM_API_KEY"))
		token := data.To
		response, err := client.Send(fcm.Message{
			Data:             data.Data,
			RegistrationIDs:  []string{token},
			ContentAvailable: true,
			Priority:         fcm.PriorityHigh,
			Notification: fcm.Notification{
				Title: "New notification from Carmind",
				Body:  "You have a new notification!",
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Println(err)
		} else {
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Status Code   :", response.StatusCode)
			fmt.Println("Success       :", response.Success)
			fmt.Println("Fail          :", response.Fail)
			fmt.Println("Canonical_ids :", response.CanonicalIDs)
			fmt.Println("Topic MsgId   :", response.MsgID)
		}
	}
}

func processNotification(token string) error {
	//Obtenemos el tiempo actual
	now := time.Now()

	//Le agregamos un minuto
	OneMinuteAgo := now.Add(-(time.Second * time.Duration(59)))

	for k, t := range listMapNotification {
		//Si el tiempo guardado en los logs, es de hace mas de un minuto, lo borramos
		if OneMinuteAgo.After(t) {
			delete(listMapNotification, k)
		}
	}

	//Verificamos si esta el token al mandar el notification
	if _, ok := listMapNotification[token]; ok {
		//Si esta, deberiamos tirar error
		return errors.New("itentelo mas tarde")
	} else {
		//Si no esta, lo dejamos proseguir y guardamos el log
		listMapNotification[token] = time.Now()
	}

	return nil
}
