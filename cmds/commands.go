package cmds

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/args"
	"github.com/VPavliashvili/wallpainter-go/cmdfactory"
	"github.com/VPavliashvili/wallpainter-go/cmds/help"
	setwallpaper "github.com/VPavliashvili/wallpainter-go/cmds/setWallpaper"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

type availableCommands struct{}

func (ac availableCommands) Get() []domain.Command {
	return available
}

func Create(input []string) (domain.Command, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}

	factory := cmdfactory.Create(availableCommands{}, args.GetParser())
	cmd, err := factory.CreateCommand(input)

	if err != nil {
		return nil, err
	}

	return cmd, nil
}

var available = []domain.Command{
	help.Create(),
    setwallpaper.Create(),
}

func validateInput(input []string) error {
	if len(input) == 0 {
		return fmt.Errorf("arguments are empty, please specify command\nor see --help for help")
	}

	args := domain.GetAllCmdArguments()
	name := input[0]

	for _, item := range args {
		if item.Flag == flags.ToFlag(name) {
			return nil
		}
	}

	return fmt.Errorf("specified command with name %v does not exist\ntype --help to view available ones", name)

}
