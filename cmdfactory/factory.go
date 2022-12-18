package cmdfactory

import (
	"github.com/VPavliashvili/slideshow-go/args/parser"
	"github.com/VPavliashvili/slideshow-go/domain"
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
	arg, _ := cf.argsParser.Parse(args)

	for _, cmd := range cmds {
		if cmd.Name() == string(arg.FlagName) {
            cmd.SetArgument(*arg)
			return cmd, nil
		}
	}

	return nil, domain.NonExistentCommandError{Argument: *arg}
}
