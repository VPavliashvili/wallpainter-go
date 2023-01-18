package operations

import (
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/folderbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/imagesbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create(arg cmds.CmdArgument) models.Operation {
	isImages := slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
		return o.Name == data.ImagesOpt
	})

	if isImages {
		return imagesbased.Create(arg)
	} else {
		return folderbased.Create(arg)
	}
}
