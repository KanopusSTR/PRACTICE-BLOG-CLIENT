package models

import (
	"client/internal/entities"
)

type WritePostRequestBody struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

type EditPostRequestBody struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

type GetPostsResponseBody struct {
	Message string           `json:"message"`
	Data    []*entities.Post `json:"data"`
}

type GetPostResponseBody struct {
	Message string        `json:"message"`
	Data    entities.Post `json:"data"`
}

type WritePostResponseBody struct {
	Message string `json:"message"`
}
