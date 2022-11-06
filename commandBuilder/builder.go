package commandbuilder

import (
	"github.com/VPavliashvili/slideshow-go/configdata"
	"golang.org/x/exp/slices"
)

func GetBuilder(data configdata.AllCommandData) Builder {
	return builder{
		allCommandsData: data,
	}
}

type Builder interface {
	Build([]string) error
}

type builder struct {
	allCommandsData configdata.AllCommandData
}

func (b builder) Build(args []string) error {
	cmdsData := b.allCommandsData.GetAllCommandData()

	if !hasOnlyOneFlagArgumentAtStartingPosition(args, cmdsData) {
		return invalidInputError{input: args}
	}

	var optArgs []string
	for i := 1; i < len(args); i++ {
		optArgs = append(optArgs, args[i])
	}
	if len(optArgs) > 0 && !configdata.IsOptName(optArgs[0]) {
		return invalidInputError{input: args}
	}

	var optNames []string
	flagName := args[0]
	for _, cmdData := range cmdsData {
		if cmdData.FlagName == configdata.Flag(flagName) {
			for _, opt := range cmdData.Opts {
				optNames = append(optNames, opt.Name)
			}
			break
		}
	}

	if len(optArgs) > 0 {
		for _, optArg := range optArgs {
			if configdata.IsOptName(optArg) && !slices.Contains(optNames, optArg) {
				return invalidInputError{input: args}
			}
		}
	}

	return nil
}

func hasOnlyOneFlagArgumentAtStartingPosition(args []string, allCmdData []configdata.CmdData) bool {
	first := args[0]
	var validFlags []configdata.Flag

	for _, item := range allCmdData {
		validFlags = append(validFlags, item.FlagName)
	}

	if !slices.Contains(validFlags, configdata.Flag(first)) {
		return false
	}

	for i := 1; i < len(args); i++ {
		arg := configdata.Flag(args[i])
		if slices.Contains(validFlags, arg) {
			return false
		}
	}

	return true
}
