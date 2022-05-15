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
