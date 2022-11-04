package commandbuilder

import (
	"github.com/VPavliashvili/slideshow-go/configdata"
	"golang.org/x/exp/slices"
)

type Option string

func GetBuilder(data configdata.AllCommandData) builder {
	return builder{
		allCommandsData: data,
	}
}

type builder struct {
	allCommandsData configdata.AllCommandData
}

func (b builder) Build(args []string) error {
	if !hasOnlyOneFlagArgumentAtStartingPosition(args, b.allCommandsData.Flags()) {
        return invalidInputError{input: args}
    }

    return nil
}

func hasOnlyOneFlagArgumentAtStartingPosition(args []string, validFlags []configdata.Flag) bool {
	first := args[0]

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
