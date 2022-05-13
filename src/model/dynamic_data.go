package model

type RecuperarContraseña struct {
	Nombre string `json:"nombre"`
	Code   string `json:"code"`
	Email  string `json:"email"`
}
