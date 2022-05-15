package manager

import (
	"github.com/Fonzeca/FastEmail/src/model"
)

type EmailManager struct {
}

func NewEmailManager() EmailManager {
	return EmailManager{}
}

func (ma *EmailManager) SendRecoverPassword(data model.RecuperarContraseña) {
	embudo := EmailChannel

	//Creamos el personalization con el to y la data dinamica
	dataAnon := map[string]interface{}{
		"nombre": data.Nombre,
		"code":   data.Code,
	}

	email := model.EmailSendGrid{
		TemplateId: "d-e932d9500c71478a8aafab7c658a6e73",
		EmailTo:    data.Email,
		Nombre:     data.Nombre,
		Data:       dataAnon,
	}

	//Enviamos el email al deamon para que se despache
	embudo <- email

}
