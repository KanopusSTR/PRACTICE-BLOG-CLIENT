package console

func (cs *consoleService) ReadLine() (string, error) {
	return cs.console.ReadLine()
}
