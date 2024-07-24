package app

import (
	"bufio"
	"client/internal/api/handlers"
	"client/internal/console"
	cs "client/internal/service/console"
	"os"
)

type serviceProvider struct {
	console console.Console

	service cs.I
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) output() console.Console {
	reader := bufio.NewReader(os.Stdin)
	s.console = console.NewConsole(reader)
	return s.console
}

func (s *serviceProvider) serviceImpl() cs.I {
	s.service = cs.New(s.output())
	return s.service
}

func (s *serviceProvider) handler(url string) handlers.I {
	return handlers.New(url, s.serviceImpl())
}
