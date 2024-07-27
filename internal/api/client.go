package api

import (
	"client/internal/api/handlers"
	"client/internal/api/readWriter"
	"client/pkg/commands"
)

type Client interface {
	Run() error
}

type client struct {
	handlers handlers.I
	rw       readWriter.I
}

func NewClient(i handlers.I, rw readWriter.I) Client {
	return &client{handlers: i, rw: rw}
}

func (c *client) Run() error {
	for {
		c.rw.Write("> ")
		line, err := c.rw.ReadLine()
		if err != nil {
			return err
		}
		switch line {
		case commands.Register:
			c.handlers.Register()
		case commands.Login:
			c.handlers.Login()
		case commands.GetProfile:
			c.handlers.GetProfile()

		case commands.WritePost:
			c.handlers.WritePost()
		case commands.GetPosts:
			c.handlers.GetPosts()
		case commands.EditPost:
			c.handlers.EditPost()
		case commands.GetPost:
			c.handlers.GetPost()
		case commands.DeletePost:
			c.handlers.DeletePost()

		case commands.WriteComment:
			c.handlers.WriteComment()
		case commands.GetComments:
			c.handlers.GetComments()
		case commands.DeleteComment:
			c.handlers.DeleteComment()

		case commands.Help:
			c.rw.WriteLine(commands.Get())
		default:
			c.rw.WriteLine(commands.Unknown)
		}
	}
}
