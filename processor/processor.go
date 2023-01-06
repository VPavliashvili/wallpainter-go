package processor

import (
	"fmt"
	"os"

	"github.com/VPavliashvili/wallpainter-go/cmds"
)

func getArgs(raw []string) []string {
	args := raw[1:]
	return args
}

func Process() {
	input := getArgs(os.Args)
	command, err := cmds.Create(input)

	if err != nil {
		fmt.Println(err)
		return
	}

    err = command.Execute()
    if err != nil {
        println(err.Error())
    }
}
