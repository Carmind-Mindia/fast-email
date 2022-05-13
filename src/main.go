package main

import (
	"fast-email/src/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	e := echo.New()

	emailChannel := make(chan model.RecuperarContraseña)
	go deamon(emailChannel)

	e.POST("/sendRecoverPassword", func(c echo.Context) error {
		data := model.RecuperarContraseña{}
		c.Bind(&data)

		emailChannel <- data

		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":5896"))
}

func deamon(channel chan model.RecuperarContraseña) {
	for {
		data := <-channel

		from := mail.NewEmail("Carmen de CarMind", "ayuda@mindiasoft.com")
		to := mail.NewEmail(data.Nombre, data.Email)

		p := mail.NewPersonalization()
		p.AddTos(to)
		p.SetDynamicTemplateData("nombre", data.Nombre)
		p.SetDynamicTemplateData("code", data.Code)

		email := mail.NewV3Mail()
		email.SetTemplateID("d-e932d9500c71478a8aafab7c658a6e73")
		email.SetFrom(from)
		email.AddPersonalizations(p)

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
