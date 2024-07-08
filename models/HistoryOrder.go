package models

import "time"

type HistoryOrder struct {
	client_name           string    `json:"client_name`
	exchange_name         string    `json:"exchange_name"`
	label                 string    `json:label`
	pair                  string    `json:pair`
	side                  string    `json:"side`
	type_                 string    `json:"type"`
	base_qty              float64   `json:"base_qty"`
	price                 float64   `json:"price"`
	algorithm_name_placed string    `json:"algorithm_name_placed"`
	lowest_sell_prc       float64   `json:"lowest_sell_prc"`
	highest_buy_prc       float64   `json:"highest_buy_prc" `
	commission_quote_qty  float64   `json:"commission_quote_qty" `
	time_placed           time.Time `json:"time_placed"`
}