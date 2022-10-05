package commands

func create(argname string) command {
	for _, cmd := range availableCommands {
		if cmd.getName() == argname {
			return cmd
		}
	}

	panic("no such command, should implement NotFoundCommand")
}
