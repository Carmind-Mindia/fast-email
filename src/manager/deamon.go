package manager

import (
	"fmt"
	"log"
	"os"

	"github.com/Fonzeca/FastEmail/src/model"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	EmailChannel chan model.EmailSendGrid
)

func Deamon() {
	//Creamos el channel
	EmailChannel = make(chan model.EmailSendGrid)

	for {

		//Esperamos un dato del canal
		data := <-EmailChannel

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
