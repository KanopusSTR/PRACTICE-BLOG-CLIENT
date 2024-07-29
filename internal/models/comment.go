package models

import (
	"client/internal/entities"
)

type WriteCommentRequestBody struct {
	Text   string `json:"text"`
	PostId int    `json:"post_id"`
}

type GetCommentsRequest struct {
	PostId int `json:"post_id"`
}

type DeleteCommentRequest struct {
	PostId    int `json:"post_id"`
	CommentId int `json:"comment_id"`
}

type GetCommentsResponse struct {
	Message string              `json:"message"`
	Data    []*entities.Comment `json:"data"`
}
