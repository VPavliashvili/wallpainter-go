package commands

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
)

type Command interface {
	String() string
	ArgNames() [][]string
	Execute() error
	SetArguments([]arguments.Argument)
}

var Available = []Command{
	&help{},
    &slideshow{},
}

func GetInvalidArgumentCommand() *invalidArgumentCommand{
    return &invalidArgumentCommand{}
}
