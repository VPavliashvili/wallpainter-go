package help

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/commands/help/builder"
)

func Create() *help{
    cmd := &help{}
    cmd.Setup(builder.Create(), arguments.GetAllArgumentInfo())
    return cmd
}

type help struct {
	Value bool
    HelpText string
    builder builder.HelpBuilder
    infos []arguments.ArgInfoPair
}

func (h *help) Setup(b builder.HelpBuilder, infos []arguments.ArgInfoPair) {
    h.builder = b
    h.infos = infos
}

func (h help) String() string {
    return "help command"
}

func (h help) ArgNames() [][]string {
	return [][]string{
		{"-h", "--help"},
	}
}

func (h *help) SetArguments(args []arguments.Argument) {
    if len(args) != 1 {
        return
    }
    name := args[0].GetName()
    if (name == "-h" || name == "--help") {
        h.Value = true
    }
}

func (h *help) Execute() error {
    var sb strings.Builder
    for _, info := range h.infos {
        sb.WriteString(h.builder.GetHelp(info.Names, info.Description))
    }

    h.HelpText = sb.String()
    fmt.Print(h.HelpText)

	return nil
}
