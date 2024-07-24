package console

import (
	"fmt"
)

func (cs *consoleService) ReadData(name string) (string, error) {
	var line string
	for {
		cs.console.Write(fmt.Sprintf("Enter %s: ", name))
		var err error
		line, err = cs.console.ReadLine()
		if err == nil {
			break
		} else {
			cs.console.WriteLine(err.Error())
		}
	}
	return line, nil
}
