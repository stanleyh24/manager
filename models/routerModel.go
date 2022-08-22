package models

type Router struct {
	ID       int32  `json:"id"`
	Ip       string `json:"ip"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Createrouter struct {
	Ip       string `json:"ip"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Routers []*Router
