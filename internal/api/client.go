package api

import (
	"client/internal/api/handlers"
	"client/pkg/commands"
)

type Client interface {
	Run() error
}

type client struct {
	handlers handlers.I
}

func NewClient(i handlers.I) Client {
	return &client{handlers: i}
}

func (c *client) Run() error {
	for {
		c.handlers.Write("> ")
		line, err := c.handlers.ReadLine()
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
			c.handlers.WriteLine(commands.Get())
		default:
			c.handlers.WriteLine(commands.Unknown)
		}
	}
}
