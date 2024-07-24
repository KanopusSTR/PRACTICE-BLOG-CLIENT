package console

import (
	"bufio"
	"fmt"
	"log"
)

func (c *console) Write(string string) error {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(c.writer)
	_, err := fmt.Fprint(c.writer, string)
	return err
}
