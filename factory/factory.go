package factory

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/commands"
	"github.com/VPavliashvili/slideshow-go/utils"
)

func GetCommand(args []arguments.Argument) commands.Command {
	cmd, err := createCommand(args)
	if err != nil {
		res := &invalidArgumentCommand{}
		res.SetArguments(args)
		return res
	}
	cmd.SetArguments(args)
	return cmd
}

func createCommand(args []arguments.Argument) (commands.Command, error) {
	if len(args) == 0 {
		return nil, emptyArgumentsError{}
	}

	err := checkForInvalidArguments(args)
	if err != nil {
		return nil, err
	}
	err = checkForDuplicateArguments(args)
	if err != nil {
		return nil, err
	}

	var result commands.Command

root:
	for _, command := range commands.Available {
		for _, names := range command.ArgNames() {
			for _, arg := range args {
				if utils.Contains(names, arg.GetName()) {
					result = command
					break root
				}
			}
		}
	}

	if result == nil {
		var names []string
		for _, arg := range args {
			names = append(names, arg.GetName())
		}
		ret := notImplementedError{}
		ret.setArgs(names)
		return nil, ret
	}

	return result, nil
}
