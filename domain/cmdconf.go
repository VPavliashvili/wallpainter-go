package domain

import (
	"strings"
)

type Command interface {
	GetArgument() Argument
	SetArgument(Argument)
	Execute() error
}

func IsOptName(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

type AvailableCommandsProvider interface {
	Get() []Command
}
