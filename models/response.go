package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"data"`
}
