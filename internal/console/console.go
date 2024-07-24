package console

import (
	"bufio"
)

type Console interface {
	ReadLine() (string, error)
	Write(string string)
	WriteLine(line string)
}

type console struct {
	reader *bufio.Reader
}

func NewConsole(reader *bufio.Reader) Console {
	return &console{reader: reader}
}
