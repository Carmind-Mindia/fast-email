package sdk

import (
	"bytes"
	"encoding/json"
	"fast-email/src/model"
	"net/http"
)

type FastEmailClient struct {
	config *Config
}

func NewEmailClient(config Config) FastEmailClient {
	return FastEmailClient{config: &config}
}

func (cli *FastEmailClient) SendRecoverPassword(email string, nombre string, code string) error {
	dataToSend := model.RecuperarContraseña{Nombre: nombre, Code: code, Email: email}

	postbody, _ := json.Marshal(dataToSend)

	url := cli.config.Url + "/sendRecoverPassword"
	_, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(postbody))
	if err != nil {
		return err
	}

	return nil
}