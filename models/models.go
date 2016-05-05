package models

type Flash struct {
	Type    string
	Message string
}

type Token struct {
	Token string `json:"token"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
