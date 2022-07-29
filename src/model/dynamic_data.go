package model

type RecuperarContraseña struct {
	Nombre string `json:"nombre"`
	Code   string `json:"code"`
	Email  string `json:"email"`
}

type EmailSendGrid struct {
	TemplateId string
	EmailTo    string
	Nombre     string
	Data       map[string]interface{}
}

type ResumenSemanalVacio struct {
	Email  string `json:"email"`
	Nombre string `json:"nombre"`
}

type ResumenSemanalLleno struct {
	Email        string            `json:"email"`
	Nombre       string            `json:"nombre"`
	Vencimientos []VencimientoView `json:"vencimientos"`
}

type VencimientoView struct {
	Documento string `json:"documento"`
	Vehiculo  string `json:"vehiculo"`
	Days      int    `json:"days"`
}

type FailureEvaluacion struct {
	Email              string `json:"email"`
	Nombre             string `json:"nombre"`
	NombreUsuario      string `json:"nombreUsuario"`
	ApellidoUsuario    string `json:"apellidoUsuario"`
	NombreVehiculo     string `json:"nombreVehiculo"`
	IdLog              int    `json:"idLog"`
	IdVehiculo         int    `json:"idVehiculo"`
	EvaluacionDateTime []int  `json:"evaluacionDateTime"`
}
