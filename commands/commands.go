package commands

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
)

type Command interface {
	String() string
	ArgNames() [][]string
	Execute() error
	SetArguments([]arguments.Argument)

	arguments() []arguments.Argument
}

var availableCommands = []Command{
	&help{},
}
