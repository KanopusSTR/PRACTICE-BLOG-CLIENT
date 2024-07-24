package handlers

import (
	"bytes"
	"client/internal/converter"
	"client/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h *handler) WriteComment() {
	id, err := h.console.ReadData("id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	text, err := h.console.ReadData("comment text")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	jsonStr := []byte(fmt.Sprintf(`{"post_id": %s, "text": "%s"}`, id, text))
	url := h.baseUrl + "/comments/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
	return
}

func (h *handler) GetComments() {
	postId, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	jsonStr := []byte(fmt.Sprintf(`{"post_id": %s}`, postId))
	url := h.baseUrl + "/comments/"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("JWT-Token", h.secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	p := &models.GetCommentsResponse{}

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
		for i, comment := range p.Data {
			h.console.WriteLine(converter.CommentToString(comment))
			if i != len(p.Data)-1 {
				h.console.WriteLine("----------")
			}
		}
	} else {
		h.console.WriteLine(p.Message)
	}
}

func (h *handler) DeleteComment() {
	postId, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	commentId, err := h.console.ReadData("comment id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	jsonStr := []byte(fmt.Sprintf(`{"post_id": %s}`, postId))
	url := h.baseUrl + "/comments/" + commentId
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
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
