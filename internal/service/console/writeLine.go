package console

func (cs *consoleService) WriteLine(line string) {
	cs.console.WriteLine(line)
}
