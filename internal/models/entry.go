package models

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
