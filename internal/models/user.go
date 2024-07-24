package models

type GetProfileResponse struct {
	Message string          `json:"message"`
	Data    ProfileResponse `json:"data"`
}

type ProfileResponse struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}
