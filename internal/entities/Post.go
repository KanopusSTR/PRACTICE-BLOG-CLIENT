package entities

import (
	"time"
)

type Post struct {
	PostId     int
	Header     string
	Body       string
	Date       time.Time
	AuthorMail string
}
