package commands

import "github.com/vpavliashvili/slideshow-go/args"

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
