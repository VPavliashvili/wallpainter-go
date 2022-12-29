package args

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
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
		Flag:        flags.Help,
		Opts:        []domain.Opt{},
		Description: "Prints this menu",
	},
	{
		Flag:        flags.SetWallpaper,
		Opts:        []domain.Opt{},
		Description: fmt.Sprintf("Sets new wallpaper\n      usage: %v /some/path/img.jpg", flags.SetWallpaper),
	},
}

func GetAll() []domain.Argument {
	return availableArgs
}
