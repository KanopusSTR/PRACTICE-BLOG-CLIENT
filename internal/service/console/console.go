package console

import "client/internal/console"

type I interface {
	ReadLine() (string, error)
	ReadData(name string) (string, error)
	ReadPassword() string
	WriteLine(string string)
	Write(string string)
}

type consoleService struct {
	console console.Console
}

func New(cons console.Console) I {
	return &consoleService{console: cons}
}
