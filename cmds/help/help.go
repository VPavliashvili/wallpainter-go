package help

import (
	"fmt"
	"strings"

	b "github.com/VPavliashvili/wallpainter-go/cmds/help/builder"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	helpdata "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/help"
	setwpdata "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/setwallpaper"
	slideshowdata "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

var available = map[flags.Flag]string {
    helpdata.Flag:helpdata.Description,
    setwpdata.Flag:setwpdata.Description,
    slideshowdata.Flag:slideshowdata.Description,
}

func Create() cmds.Command {
	return &help{}
}

type help struct {
}

func (h help) Name() string {
	return flags.Help
}

func (h *help) SetArgument(a cmds.CmdArgument) {}

func (h help) Execute() error {
	builder := b.Create()
	var sb strings.Builder

	for key, val := range available {
		sb.WriteString(builder.GetHelp(key, val))
	}

	result := sb.String()
	fmt.Print(result)

	return nil
}
