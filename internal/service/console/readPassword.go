package console

import (
	"gitlab.com/david_mbuvi/go_asterisks"
	"os"
)

func (cs *consoleService) ReadPassword() (string, error) {
	for {
		err := cs.console.Write("Enter password: ")
		if err != nil {
			return "", err
		}
		password, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		if err != nil {
			err2 := cs.console.WriteLine(err.Error())
			if err2 != nil {
				return "", err2
			}
		} else {
			return string(password), nil
		}
	}
}
