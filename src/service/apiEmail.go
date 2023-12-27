package service

import (
	"net/http"

	"github.com/Carmind-Mindia/fast-email/src/manager"
	"github.com/Carmind-Mindia/fast-email/src/model"
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
