package commands

var availableCommands []command

type command interface {
	getName() string
	Execute() error
}

func setup() {
	helpcmd := help{
		Arg:  args.Help,
		text: "help text placeholder",
	}

	availableCommands = append(availableCommands, helpcmd)
}
