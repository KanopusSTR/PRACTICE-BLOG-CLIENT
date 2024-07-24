package console

import (
	"bufio"
	"fmt"
	"log"
)

func (c *console) WriteLine(line string) error {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}(c.writer)
	_, err := fmt.Fprintln(c.writer, line)
	return err
}
