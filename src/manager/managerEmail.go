package manager

import "github.com/Fonzeca/FastEmail/src/model"

type EmailManager struct {
}

func NewEmailManager() EmailManager {
	return EmailManager{}
}

func (ma *EmailManager) SendRecoverPassword(data model.RecuperarContraseña) {
	embudo := EmailChannel

	dataAnon := map[string]interface{}{
		"code": data.Code,
	}

	email := model.EmailTemplate{
		TemplateId: "6585f6008759392fbc0d73be",
		EmailTo:    data.Email,
		Data:       dataAnon,
	}

	//Enviamos el email al deamon para que se despache
	embudo <- email

}

// func (ma *EmailManager) SendDocsCloseToExpire(data model.ResumenSemanalLleno) {
// 	embudo := EmailChannel

// 	//Creamos el personalization con el to y la data dinamica
// 	resumenSemanalLleno := map[string]interface{}{
// 		"nombre":       data.Nombre,
// 		"vencimientos": data.Vencimientos,
// 	}

// 	email := model.EmailSendGrid{
// 		TemplateId: "d-1dd035cceb9b4d23b4af4867af3956da",
// 		EmailTo:    data.Email,
// 		Nombre:     data.Nombre,
// 		Data:       resumenSemanalLleno,
// 	}

// 	//Enviamos el email al deamon para que se despache
// 	embudo <- email

// }

// func (ma *EmailManager) SendNoneDocsCloseToExpire(data model.ResumenSemanalVacio) {
// 	embudo := EmailChannel

// 	//Creamos el personalization con el to y la data dinamica
// 	resumenSemanalVacio := map[string]interface{}{
// 		"nombre": data.Nombre,
// 	}

// 	email := model.EmailSendGrid{
// 		TemplateId: "d-beee6c66fe054159ac7e1e5b0e68f911",
// 		EmailTo:    data.Email,
// 		Nombre:     data.Nombre,
// 		Data:       resumenSemanalVacio,
// 	}

// 	//Enviamos el email al deamon para que se despache
// 	embudo <- email

// }

// func (ma *EmailManager) SendFailureEvaluacion(data model.FailureEvaluacion) {
// 	embudo := EmailChannel

// 	//Creamos el personalization con el to y la data dinamica
// 	failureEvaluacion := map[string]interface{}{
// 		"nombreUsuario":      data.NombreUsuario,
// 		"apellidoUsuario":    data.ApellidoUsuario,
// 		"nombreVehiculo":     data.NombreVehiculo,
// 		"idLog":              data.IdLog,
// 		"idVehiculo":         data.IdVehiculo,
// 		"evaluacionDateTime": data.EvaluacionDateTime,
// 	}

// 	email := model.EmailSendGrid{
// 		TemplateId: "d-c1bb791cbd8c4fdeb2067993f9c14597",
// 		EmailTo:    data.Email,
// 		Nombre:     data.Nombre,
// 		Data:       failureEvaluacion,
// 	}

// 	//Enviamos el email al deamon para que se despache
// 	embudo <- email
// }
