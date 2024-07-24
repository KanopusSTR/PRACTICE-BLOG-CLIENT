package console

import (
	"bufio"
)

type Console interface {
	ReadLine() (string, error)
	Write(string string) error
	WriteLine(line string) error
}

type console struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewConsole(reader *bufio.Reader, writer *bufio.Writer) Console {
	return &console{reader: reader, writer: writer}
}
