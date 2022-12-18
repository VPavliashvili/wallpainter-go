package cmdfactory

import (
	"github.com/VPavliashvili/slideshow-go/args/parser"
	"github.com/VPavliashvili/slideshow-go/domain"
)

var availableCommands domain.AvailableCommandsProvider
var argsparser parser.Parser

func Setup(commandsprovider domain.AvailableCommandsProvider, parser parser.Parser) {
	availableCommands = commandsprovider
	argsparser = parser
}

func Create(args []string) (domain.Command, error) {
	cmds := availableCommands.Get()
	arg, _ := argsparser.Parse(args)

	for _, cmd := range cmds {
		if cmd.GetArgument().FlagName == arg.FlagName {
			return cmd, nil
		}
	}

	return nil, domain.NonExistentCommandError{Argument: *arg}
}
