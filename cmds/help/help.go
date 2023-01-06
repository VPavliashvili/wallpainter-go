package help

import (
	"fmt"
	"strings"

	b "github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

var setwprpdesc = fmt.Sprintf(`Sets new wallpaper
      usage: %v /some/path/img.jpg
      options: --scaling {max, fill, center, tile, scale}
      scale is default`, flags.SetWallpaper)

var availableArgs = []domain.CmdArgument{
	{
		Flag:        flags.Help,
		Description: "Prints this menu",
	},
	{
		Flag:        flags.SetWallpaper,
		Description: setwprpdesc,
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
	builder := b.Create()
	var sb strings.Builder

	for _, arg := range h.predefined {
		sb.WriteString(builder.GetHelp(arg))
	}

	result := sb.String()
	fmt.Print(result)

	return nil
}
