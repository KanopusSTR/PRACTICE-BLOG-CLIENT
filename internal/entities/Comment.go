package entities

import (
	"time"
)

type Comment struct {
	CommentId  int
	Text       string
	Date       time.Time
	AuthorMail string
	PostId     int
}
