package cmdfactory

import (
	"errors"

	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/domain"
)

func Create(commandsprovider domain.AvailableCommandsProvider, parser parser.Parser) CommandFactory {
	return factory{
		availableCommands: commandsprovider,
		argsParser:        parser,
	}
}

type CommandFactory interface {
	CreateCommand([]string) (domain.Command, error)
}

type factory struct {
	availableCommands domain.AvailableCommandsProvider
	argsParser        parser.Parser
}

func (cf factory) CreateCommand(args []string) (domain.Command, error) {
	cmds := cf.availableCommands.Get()
	arg, err := cf.argsParser.Parse(args)

	if err != nil {
		return nil, err
	}

	for _, cmd := range cmds {
		if cmd.Name() == string(arg.Flag) {
			cmd.SetArgument(*arg)
			return cmd, nil
		}
	}

	return nil, errors.New("this error should have been handled before this line")
}
