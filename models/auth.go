package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Key string `json:"key"`
}
