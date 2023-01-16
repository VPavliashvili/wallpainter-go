package args

import (
	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/args/parser/helpopts"
	setwallpaperopts "github.com/VPavliashvili/wallpainter-go/args/parser/setWallpaperopts"
	"github.com/VPavliashvili/wallpainter-go/args/parser/slideshowopts"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func GetParser() parser.Parser {
	return parser.Create(provider{})
}

type provider struct{}

func (p provider) Get() []domain.RawArgument {
	return rawavailable
}

var rawavailable = []domain.RawArgument{
	{
		Flag:       flags.Help,
		OptsParser: helpopts.Create(),
	},
	{
		Flag:       flags.SetWallpaper,
		OptsParser: setwallpaperopts.Create(),
	},
	{
		Flag:       flags.RunSlideShow,
		OptsParser: slideshowopts.Create(),
	},
}
