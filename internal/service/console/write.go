package console

import "log"

func (cs *consoleService) Write(string string) {
	err := cs.console.Write(string)
	if err != nil {
		log.Fatal(err)
		return
	}
}
