package help

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

var availableArgs = []domain.CmdArgument{
	{
		Flag:        flags.Help,
		Description: "Prints this menu",
	},
	{
		Flag:        flags.SetWallpaper,
		Description: fmt.Sprintf("Sets new wallpaper\n      usage: %v /some/path/img.jpg", flags.SetWallpaper),
	},
}

func Create() domain.Command {
	return &help{
		predefined: availableArgs,
	}
}

type help struct {
	predefined []domain.CmdArgument
}

func (h help) Name() string {
	return flags.Help
}

func (h *help) SetArgument(a domain.CmdArgument) {}

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
