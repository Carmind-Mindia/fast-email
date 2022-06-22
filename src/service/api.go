package service

import (
	"net/http"

	"github.com/Fonzeca/FastEmail/src/manager"
	"github.com/Fonzeca/FastEmail/src/model"
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

func (api *ApiEmail) SendDocsCloseToExpire(c echo.Context) error {
	data := model.ResumenSemanalLleno{}
	c.Bind(&data)

	api.manager.SendDocsCloseToExpire(data)

	return c.NoContent(http.StatusOK)
}

func (api *ApiEmail) SendNoneDocsCloseToExpire(c echo.Context) error {
	data := model.ResumenSemanalVacio{}
	c.Bind(&data)

	api.manager.SendNoneDocsCloseToExpire(data)

	return c.NoContent(http.StatusOK)
}
