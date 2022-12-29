package help

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/args"
	"github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func Create() domain.Command {
	allargs := args.GetAll()
	//arg := args.GetByName(flag)

	return &help{
		//argument:   arg,
		predefined: allargs,
	}
}

type help struct {
	argument   domain.Argument
	predefined []domain.Argument
}

func (h help) Name() string {
	return flags.Help
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
