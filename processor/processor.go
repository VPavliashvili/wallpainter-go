package processor

import (
	"os"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/commands"
)

func getArgs(raw []string) []string {
    args := raw[1:]
    return args
}

func Process() {
    args, _ := arguments.GetArguments(getArgs(os.Args))
    cmd := commands.GetCommand(args)

    cmd.Execute()
}
