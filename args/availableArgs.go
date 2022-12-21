package args

import (
	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/domain"
)

type argumentsProvider struct{}

func (ap argumentsProvider) Get() []domain.Argument {
	return availableArgs
}

func GetParser() parser.Parser {
	return parser.Create(argumentsProvider{})
}

var availableArgs = []domain.Argument{
	{
		FlagName: "--help",
		Opts:     []domain.Opt{},
        Description: "Prints this menu",
	},
}

func GetAll() []domain.Argument{
    return availableArgs
}