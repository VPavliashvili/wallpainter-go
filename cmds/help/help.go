package help

import (
	"fmt"
	"strings"

	b "github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

var setwprpdesc = fmt.Sprintf(`Sets new wallpaper
      usage: %v /some/path/img.jpg
      options: --scaling {max, fill, center, tile, scale}
      scale is default`, flags.SetWallpaper)

var availableArgs = []cmds.CmdArgument{
	{
		Flag:        flags.Help,
		Description: "Prints this menu",
	},
	{
		Flag:        flags.SetWallpaper,
		Description: setwprpdesc,
	},
}

func Create() cmds.Command {
	return &help{
		predefined: availableArgs,
	}
}

type help struct {
	predefined []cmds.CmdArgument
}

func (h help) Name() string {
	return flags.Help
}

func (h *help) SetArgument(a cmds.CmdArgument) {}

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
