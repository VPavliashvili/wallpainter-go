package commands

import (
	"fmt"
)

type help struct {
	Arg  args.Argument[bool]
	text string
}

func (cmd help) getName() string {
	return cmd.Arg.Names
}

func (cmd help) Execute() error {
	fmt.Println(cmd.text)

	return nil
}
