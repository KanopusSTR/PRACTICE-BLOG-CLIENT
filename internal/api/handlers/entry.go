package handlers

import (
	"bytes"
	"client/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *handler) Register() {
	username, err := h.console.ReadData("username")

	mail, err := h.console.ReadData("user mail")

	password := h.console.ReadPassword()

	url := h.baseUrl + "/register"

	jsonStr := []byte(fmt.Sprintf(`{"name": "%s", "mail": "%s", "password": "%s"}`, username, mail, password))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.ResultResponseBody{}

	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		h.console.WriteLine(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	h.console.WriteLine(p.Message)
}

func (h *handler) Login() {
	mail, err := h.console.ReadData("user mail")

	password := h.console.ReadPassword()

	url := h.baseUrl + "/login"

	jsonStr := []byte(fmt.Sprintf(`{"mail": "%s", "password": "%s"}`, mail, password))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.LoginResponseBody{}

	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	h.console.WriteLine(p.Message)
	if p.Message == "success" {
		h.secretKey = p.Data
	}
}

func (h *handler) GetProfile() {
	mail, err := h.console.ReadData("user mail")

	url := h.baseUrl + "/users/" + mail
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.GetProfileResponse{}

	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			h.console.WriteLine(err.Error())
		}
	}(resp.Body)
	if p.Message == "success" {
		h.console.WriteLine(fmt.Sprintf("username: %s, profile mail: %s", p.Data.Name, p.Data.Mail))
	} else {
		h.console.WriteLine(p.Message)
	}
}
