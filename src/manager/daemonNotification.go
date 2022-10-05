package manager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Fonzeca/FastEmail/src/model"
	"github.com/spf13/viper"
	"google.golang.org/api/fcm/v1"
	"google.golang.org/api/option"
)

var (
	NotificationChannel chan model.CarmindNotification
)

var listMapNotification = make(map[string]time.Time)

func DeamonNotification() {
	//Creamos el channel
	NotificationChannel = make(chan model.CarmindNotification)

	ctx := context.Background()

	opt := option.WithCredentialsFile(viper.GetString("FIREBASE_CREDENTIAL_FILE"))
	fcmService, err := fcm.NewService(ctx, opt)

	if err != nil {
		log.Fatal(err)
	}

	for {

		//Esperamos un dato del canal
		data := <-NotificationChannel

		for _, token := range data.To {
			err := processNotification(token)
			if err != nil {
				//TODO: logeamos el error
				fmt.Print(err)
				continue
			}
		}

		for _, token := range data.To {
			sendMessageRequest := &fcm.SendMessageRequest{
				Message: &fcm.Message{
					Token: token,
					Notification: &fcm.Notification{
						Title: data.Data["Title"].(string),
						Body:  data.Data["Message"].(string),
					},
				},
			}
			projectMessage := fcmService.Projects.Messages.Send("projects/carmind-46f12", sendMessageRequest)
			_, err := projectMessage.Do()
			if err != nil {
				fmt.Print(err.Error())
				continue
			}
		}
	}
}

func processNotification(tokens string) error {
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
	if _, ok := listMapNotification[tokens]; ok {
		//Si esta, deberiamos tirar error
		return errors.New("itentelo mas tarde")
	} else {
		//Si no esta, lo dejamos proseguir y guardamos el log
		listMapNotification[tokens] = time.Now()
	}

	return nil
}
