package readWriter

import "client/internal/service/console"

type I interface {
	ReadLine() (string, error)
	WriteLine(strings ...string)
	Write(string string)
}

type readWriter struct {
	console console.I
}

func New(console console.I) I {
	return &readWriter{console: console}
}

func (rw *readWriter) ReadLine() (string, error) {
	return rw.console.ReadLine()
}

func (rw *readWriter) WriteLine(strings ...string) {
	for _, s := range strings {
		rw.console.WriteLine(s)
	}
}

func (rw *readWriter) Write(string string) {
	rw.console.Write(string)
}
