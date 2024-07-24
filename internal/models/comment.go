package models

import (
	"client/internal/entities"
)

type WriteCommentRequestBody struct {
	Text   string `json:"text"`
	PostId int    `json:"post_id"`
}

type GetCommentsRequestBody struct {
	PostId   int                 `json:"post_id"`
	Comments []*entities.Comment `json:"comments"`
}

type GetCommentsResponse struct {
	Message string              `json:"message"`
	Data    []*entities.Comment `json:"data"`
}
