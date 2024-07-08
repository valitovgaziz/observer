package models

import (

)

type Client struct {
	client_name   string `json:"client_name"`
	exchange_name string `json:"exchange_name"`
	label         string `json:"label"`
	pair          string `json:"pair"`
}