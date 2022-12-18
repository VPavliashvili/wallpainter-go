package cmdfactory

import (
	"github.com/VPavliashvili/slideshow-go/domain"
)

var availableCommands domain.AvailableCommandsProvider

func Setup(provider domain.AvailableCommandsProvider) {
	availableCommands = provider
}

func Create(arg domain.Argument) (domain.Command, error) {
	cmds := availableCommands.Get()

	for _, cmd := range cmds {
		if cmd.GetArgument().FlagName == arg.FlagName {
			return cmd, nil
		}
	}

	return nil, domain.NonExistentCommandError{Argument: arg}
}
