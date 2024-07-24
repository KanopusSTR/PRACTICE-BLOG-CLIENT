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

	ReadLine() (string, error)
	WriteLine(strings ...string)
	Write(string string)
}

type handler struct {
	baseUrl   string
	secretKey string

	console console.I
}

func New(baseUrl string, i console.I) I {
	return &handler{baseUrl: baseUrl, secretKey: "", console: i}
}

func (h *handler) ReadLine() (string, error) {
	return h.console.ReadLine()
}

func (h *handler) WriteLine(strings ...string) {
	for _, s := range strings {
		h.console.WriteLine(s)
	}
}

func (h *handler) Write(string string) {
	h.console.Write(string)
}
