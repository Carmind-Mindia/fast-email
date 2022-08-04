package model

type CarmindNotification struct {
	Prioity string
	To      string
	Data    map[string]interface{}
}

type SimpleNotification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	To      string `json:"to"`
}
