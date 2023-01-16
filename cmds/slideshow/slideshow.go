package slideshow

import (
	"errors"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func Create() cmds.Command {
    return runslideshow{}
}

type runslideshow struct {
}

func (r runslideshow) SetArgument(arg cmds.CmdArgument) {}
func (r runslideshow) Execute() error {
    return errors.New("not implemented yet")
}
func (r runslideshow) Name() string {
    return flags.RunSlideShow
}
