package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Fonzeca/FastEmail/src/model"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	EmailChannel chan model.EmailSendGrid
)

var listMapEmail = make(map[string]time.Time)

func DeamonEmail() {
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
			jsonStr, err := json.MarshalIndent(email, "", "\t")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Status code: " + strconv.Itoa(response.StatusCode))
			fmt.Println("Response body:" + response.Body)
			fmt.Println("Request body:\n" + string(jsonStr))
		}
	}
}

func processEmail(email string) error {
	//Obtenemos el tiempo actual
	now := time.Now()

	//Le agregamos un minuto
	OneMinuteAgo := now.Add(-(time.Second * time.Duration(59)))

	for k, t := range listMapEmail {
		//Si el tiempo guardado en los logs, es de hace mas de un minuto, lo borramos
		if OneMinuteAgo.After(t) {
			delete(listMapEmail, k)
		}
	}

	//Verificamos si esta el correo a mandar el email
	if _, ok := listMapEmail[email]; ok {
		//Si esta, deberiamos tirar error
		return errors.New("intentelo mas tarde")
	} else {
		//Si no esta, lo dejamos proseguir y guardamos el log
		listMapEmail[email] = time.Now()
	}

	return nil
}
