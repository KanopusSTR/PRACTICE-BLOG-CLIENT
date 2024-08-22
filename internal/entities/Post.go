package entities

import (
	"time"
)

type Post struct {
	Id         int
	Header     string
	Body       string
	Date       time.Time
	AuthorMail string
}
