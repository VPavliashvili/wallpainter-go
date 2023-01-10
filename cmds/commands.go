package cmds

import (
	"github.com/VPavliashvili/wallpainter-go/args"
	"github.com/VPavliashvili/wallpainter-go/cmdfactory"
	"github.com/VPavliashvili/wallpainter-go/cmds/help"
	setwallpaper "github.com/VPavliashvili/wallpainter-go/cmds/setWallpaper"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
)

type availableCommands struct{}

func (ac availableCommands) Get() []cmds.Command {
	return available
}

func Create(input []string) (cmds.Command, error) {
	factory := cmdfactory.Create(availableCommands{}, args.GetParser())
	cmd, err := factory.CreateCommand(input)

	if err != nil {
		return nil, err
	}

	return cmd, nil
}

var available = []cmds.Command{
	help.Create(),
    setwallpaper.Create(),
}

