package processor

import (
	"fmt"
	"os"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/factory"
)

func getArgs(raw []string) []string {
    args := raw[1:]
    return args
}

func Process() {
    args, _ := arguments.GetArguments(getArgs(os.Args))
    cmd := factory.GetCommand(args)

    err := cmd.Execute()
    if err != nil {
        fmt.Println(err)
    }
}
