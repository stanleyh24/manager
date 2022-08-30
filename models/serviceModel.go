package models

type Service struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Rate  string `json:"rate"`
}

type CreateServiceParams struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Rate  string `json:"rate"`
}
