package console

import (
	"gitlab.com/david_mbuvi/go_asterisks"
	"os"
)

func (cs *consoleService) ReadPassword() string {
	for {
		cs.console.Write("Enter password: ")
		password, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		if err != nil {
			cs.console.WriteLine(err.Error())
		} else {
			return string(password)
		}
	}
}
