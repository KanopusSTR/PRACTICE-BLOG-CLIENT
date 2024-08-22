package handlers

import (
	"bytes"
	"client/internal/converter"
	"client/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (h *handler) WritePost() {
	header, err := h.console.ReadData("post header")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	body, err := h.console.ReadData("post body")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	st := models.WritePostRequestBody{
		Header: header,
		Body:   body,
	}
	jsonStr, err := json.Marshal(st)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	url := h.baseUrl + "/posts/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.WritePostResponseBody{}

	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			h.console.WriteLine(err.Error())
		}
	}(resp.Body)
	h.console.WriteLine(p.Message)
}

func (h *handler) GetPosts() {
	url := h.baseUrl + "/posts/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.GetPostsResponseBody{}

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
		for i, post := range p.Data {
			h.console.WriteLine(converter.PostToString(post))
			if i != len(p.Data)-1 {
				h.console.WriteLine("----------")
			}
		}
	} else {
		h.console.WriteLine(p.Message)
	}
}

func (h *handler) EditPost() {
	id, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	header, err := h.console.ReadData("post header")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	body, err := h.console.ReadData("post body")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	st := models.EditPostRequestBody{
		Header: header,
		Body:   body,
	}
	jsonStr, err := json.Marshal(st)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	url := h.baseUrl + "/posts/" + id
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.ResultResponseBody{}

	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			h.console.WriteLine(err.Error())
		}
	}(resp.Body)
	h.console.WriteLine(p.Message)
}

func (h *handler) GetPost() {
	id, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	url := h.baseUrl + "/posts/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.GetPostResponseBody{}

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
		h.console.WriteLine(converter.PostToString(&p.Data))
	} else {
		h.console.WriteLine(p.Message)
	}
}

func (h *handler) DeletePost() {
	ids, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	_, err = strconv.Atoi(ids)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	url := h.baseUrl + "/posts/" + ids
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.ResultResponseBody{}

	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			h.console.WriteLine(err.Error())
		}
	}(resp.Body)
	h.console.WriteLine(p.Message)
}
