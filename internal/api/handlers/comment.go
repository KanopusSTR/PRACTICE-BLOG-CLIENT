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

func (h *handler) WriteComment() {
	ids, err := h.console.ReadData("id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	id, err := strconv.Atoi(ids)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	text, err := h.console.ReadData("comment text")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	st := models.WriteCommentRequestBody{PostId: id, Text: text}
	jsonStr, err := json.Marshal(st)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	url := h.baseUrl + "/comments/"
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

func (h *handler) GetComments() {
	postId, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	id, err := strconv.Atoi(postId)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	st := models.GetCommentsRequest{PostId: id}
	jsonStr, err := json.Marshal(st)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	url := h.baseUrl + "/comments/"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
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
	postIds, err := h.console.ReadData("post id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	postId, err := strconv.Atoi(postIds)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	commentIds, err := h.console.ReadData("comment id")
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	commentId, err := strconv.Atoi(commentIds)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}

	st := models.DeleteCommentRequest{PostId: postId, CommentId: commentId}
	jsonStr, err := json.Marshal(st)
	if err != nil {
		h.console.WriteLine(err.Error())
		return
	}
	url := h.baseUrl + "/comments/" + commentIds
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
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
