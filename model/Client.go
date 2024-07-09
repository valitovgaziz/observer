package model

type Client struct {
	Client_name   string `json:"client_name"`
	Exchange_name string `json:"exchange_name"`
	Label         string `json:"label"`
	Pair          string `json:"pair"`
}