package slideshow

import (
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func Create() cmds.Command {
	return &runslideshow{}
}

type runslideshow struct {
	operation models.Operation
}

func (r *runslideshow) SetArgument(arg cmds.CmdArgument) {
	r.operation = operations.Create(arg)
}

func (r runslideshow) Execute() error {
	return r.operation.Execute()
}
func (r runslideshow) Name() string {
	return flags.RunSlideShow
}
