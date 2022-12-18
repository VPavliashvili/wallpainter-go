package cmds

import (
	"fmt"

	"github.com/VPavliashvili/slideshow-go/args"
	"github.com/VPavliashvili/slideshow-go/cmdfactory"
	"github.com/VPavliashvili/slideshow-go/cmds/help"
	"github.com/VPavliashvili/slideshow-go/domain"
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
}

func validateInput(input []string) error {
    if len(input)==0{
        return fmt.Errorf("arguments are empty, please specify command\nor see --help for help")
    }

	args := args.GetAll()
	name := input[0]

	for _, item := range args {
		if item.FlagName == domain.Flag(name) {
			return nil
		}
	}

	return fmt.Errorf("specified command with name %v does not exist\ntype --help to view available ones", name)

}
