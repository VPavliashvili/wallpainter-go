package help

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/args"
	"github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain"
)

func Create() domain.Command {
	allargs := args.GetAll()

	var arg domain.Argument
	for _, item := range allargs {
		if item.FlagName == "--help" {
			arg = item
			break
		}
	}

	return &help{
		argument:   arg,
		predefined: allargs,
	}
}

type help struct {
	argument   domain.Argument
	predefined []domain.Argument
}

func (h help) GetArgument() domain.Argument {
	return h.argument
}

func (h help) Name() string {
	return "--help"
}

func (h *help) SetArgument(a domain.Argument) {
	h.argument = a
}

func (h help) Execute() error {
	builder := builder.Create()
	var sb strings.Builder

	for _, arg := range h.predefined {
		sb.WriteString(builder.GetHelp(arg))
	}

    result := sb.String()
	fmt.Print(result)

	return nil
}
