package factory

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

type invalidArgumentCommand struct {
	input []string
    err error
}

func (i invalidArgumentCommand) String() string {
	return strings.Join(i.input, ", ")
}

func (i invalidArgumentCommand) ArgNames() [][]string {
	return nil
}

func (i invalidArgumentCommand) Execute() error {
    msg := fmt.Sprintf("wrong input specified %v\nerror: %v\nsee --help/-h", i.String(), i.err)
    fmt.Println(msg)
    return nil
}

func (i *invalidArgumentCommand) SetArguments(args []arguments.Argument) {
    var names []string
    for _, arg := range args {
        names = append(names, arg.GetName())
    }
    i.input = names
}
