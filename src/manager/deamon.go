package manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Fonzeca/FastEmail/src/model"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	EmailChannel chan model.EmailSendGrid
)

var listMap = make(map[string]time.Time)

func Deamon() {
	//Creamos el channel
	EmailChannel = make(chan model.EmailSendGrid)

	for {

		//Esperamos un dato del canal
		data := <-EmailChannel

		err := processEmail(data.EmailTo)
		if err != nil {
			//TODO: logeamos el error
			fmt.Print(err)
			continue
		}

		from := mail.NewEmail("Carmen de CarMind", "ayuda@mindiasoft.com")

		pers := mail.NewPersonalization()
		pers.AddTos(mail.NewEmail(data.Nombre, data.EmailTo))
		pers.DynamicTemplateData = data.Data

		email := mail.NewV3Mail()
		email.SetTemplateID(data.TemplateId)
		email.SetFrom(from)
		email.AddPersonalizations(pers)

		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(email)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}
}

func processEmail(email string) error {
	//Obtenemos el tiempo actual
	now := time.Now()

	//Le agregamos un minuto
	OneMinuteAgo := now.Add(-(time.Second * time.Duration(59)))

	for k, t := range listMap {
		//Si el tiempo guardado en los logs, es de hace mas de un minuto, lo borramos
		if OneMinuteAgo.After(t) {
			delete(listMap, k)
		}
	}

	//Verificamos si esta el correo a mandar el email
	if _, ok := listMap[email]; ok {
		//Si esta, deberiamos tirar error
		return errors.New("Intentelo mas tarde")
	} else {
		//Si no esta, lo dejamos proseguir y guardamos el log
		listMap[email] = time.Now()
	}

	return nil
}
