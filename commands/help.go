package commands

import (
	"fmt"

	"github.com/vpavliashvili/slideshow-go/args"
)

type help struct {
	Arg  args.Argument[bool]
	text string
}

func (cmd help) getName() string {
	return cmd.Arg.Name
}

func (cmd help) Execute() error {
	fmt.Println(cmd.text)

	return nil
}
