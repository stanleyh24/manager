package models

type Client struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Last_name    string `json:"last_name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Payment_date int    `json:"payment_date"`
	Service_id   int    `json:"service_id"`
	Router_id    int    `json:"router_id"`
}

type CreateClientParams struct {
	Name         string `json:"name"`
	Last_name    string `json:"last_name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Payment_date int    `json:"payment_date"`
	Service_id   int    `json:"service_id"`
	Router_id    int    `json:"router_id"`
}
