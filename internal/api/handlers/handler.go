package handlers

import (
	"client/internal/service/console"
)

type I interface {
	Register()
	Login()
	GetProfile()

	WritePost()
	GetPosts()
	EditPost()
	GetPost()
	DeletePost()

	WriteComment()
	GetComments()
	DeleteComment()
}

type handler struct {
	baseUrl   string
	secretKey string

	console console.I
}

func New(baseUrl string, i console.I) I {
	return &handler{baseUrl: baseUrl, secretKey: "", console: i}
}
