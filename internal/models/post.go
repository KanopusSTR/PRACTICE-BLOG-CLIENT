package models

import (
	"client/internal/entities"
)

type PostRequestBody struct {
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
	Data    int    `json:"data"`
}
