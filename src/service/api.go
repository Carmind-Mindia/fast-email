package service

import (
	"fast-email/src/manager"
	"fast-email/src/model"
	"net/http"

	"github.com/labstack/echo"
)

type ApiEmail struct {
	manager manager.EmailManager
}

func NewApiEmail() ApiEmail {
	m := manager.NewEmailManager()
	return ApiEmail{manager: m}
}

func (api *ApiEmail) SendRecoverPassword(c echo.Context) error {
	data := model.RecuperarContraseña{}
	c.Bind(&data)

	api.manager.SendRecoverPassword(data)

	return c.NoContent(http.StatusOK)
}
