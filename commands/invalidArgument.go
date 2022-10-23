package commands

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
)

type invalidArgumentCommand struct{
    input []string
}

func (i invalidArgumentCommand) String() string {
    return ""
}

func (i invalidArgumentCommand) ArgNames() [][]string{
    return nil
}

func (i invalidArgumentCommand) Execute() error {
    return nil
}

func (i *invalidArgumentCommand) SetArguments(args []arguments.Argument) {
}

func (i invalidArgumentCommand) arguments() []arguments.Argument {
    return nil
}
