package models

type Router struct {
	ID       int64  `json:"id"`
	Ip       string `json:"ip"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateRouterParams struct {
	Ip       string `json:"ip"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
