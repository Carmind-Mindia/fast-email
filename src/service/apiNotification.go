package service

import (
	"net/http"

	"github.com/Fonzeca/FastEmail/src/manager"
	"github.com/Fonzeca/FastEmail/src/model"
	"github.com/labstack/echo"
)

type ApiNotification struct {
	manager manager.NotificationManager
}

func NewApiNotification() ApiNotification {
	m := manager.NewNotificationManager()
	return ApiNotification{manager: m}
}

func (api *ApiNotification) SendNotificationToCarmind(c echo.Context) error {
	data := model.SimpleNotification{}
	c.Bind(&data)

	api.manager.SendNotificationToCarmind(data)

	return c.NoContent(http.StatusOK)
}
