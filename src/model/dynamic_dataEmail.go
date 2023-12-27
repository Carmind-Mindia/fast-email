package model

type RecuperarContraseña struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

type EmailTemplate struct {
	TemplateId string
	EmailTo    string
	Data       map[string]interface{}
}

type EnvialoSimpleApiBody struct {
	From          string                 `json:"from"`
	To            []string               `json:"to"`
	TemplateId    string                 `json:"templateID"`
	Subject       string                 `json:"subject"`
	Substitutions map[string]interface{} `json:"substitutions"`
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
	EvaluacionDateTime string `json:"evaluacionDateTime"`
}

type ZoneNotification struct {
	Imei         string   `json:"imei"`
	ZoneName     string   `json:"zone_name"`
	ZoneID       int      `json:"zone_id"`
	EventType    string   `json:"event_type"`
	VehiculoId   int      `json:"vehiculo_id"`
	VehiculoName string   `json:"vehiculo_name"`
	Emails       []string `json:"emails"`
	FCMTokens    []string `json:"fcmtokens"`
}
