package app

import (
	"client/internal/api/handlers"
	"client/internal/api/readWriter"
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
	s.console = console.NewConsole(os.Stdin, os.Stdout, nil, nil)
	return s.console
}

func (s *serviceProvider) serviceImpl() cs.I {
	s.service = cs.New(s.output())
	return s.service
}

func (s *serviceProvider) handler(url string) handlers.I {
	return handlers.New(url, s.serviceImpl())
}

func (s *serviceProvider) readWriter() readWriter.I {
	return readWriter.New(s.service)
}
