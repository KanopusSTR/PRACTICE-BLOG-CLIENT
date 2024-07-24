package console

import (
	"fmt"
)

func (cs *consoleService) ReadData(name string) (string, error) {
	var line string
	for {
		err := cs.console.Write(fmt.Sprintf("Enter %s: ", name))
		if err != nil {
			return "", err
		}
		line, err = cs.console.ReadLine()
		if err == nil {
			break
		} else {
			err = cs.console.WriteLine(err.Error())
			if err != nil {
				return "", err
			}
		}
	}
	return line, nil
}
