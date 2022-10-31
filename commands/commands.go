package commands

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/commands/help"
	sw "github.com/VPavliashvili/slideshow-go/commands/setWallpaper"
	ss "github.com/VPavliashvili/slideshow-go/commands/slideshow"
)

type Command interface {
	String() string
	ArgNames() [][]string
	Execute() error
	SetArguments([]arguments.Argument)
}

var Available = []Command{
	help.Create(),
    ss.Create(),
    sw.Create(),
}
