package console

import (
	"bufio"
	"os"
)

type Console interface {
	ReadLine() (string, error)
	Write(string string) error
	WriteLine(line string) error

	GetFileReader() *os.File
	GetFileWriter() *os.File
}

type console struct {
	reader *bufio.Reader
	writer *bufio.Writer

	fileR *os.File
	fileW *os.File
}

func NewConsole(fileR *os.File, fileW *os.File, reader *bufio.Reader, writer *bufio.Writer) Console {
	if reader == nil {
		reader = bufio.NewReader(fileR)
	}
	if writer == nil {
		writer = bufio.NewWriter(fileW)
	}
	return &console{reader: reader, writer: writer, fileR: fileR, fileW: fileW}
}

func (c *console) GetFileReader() *os.File {
	return c.fileR
}

func (c *console) GetFileWriter() *os.File {
	return c.fileW
}
