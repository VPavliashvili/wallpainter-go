package operations

import (
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/folderbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/helpbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/imagesbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/listimages"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create(arg cmds.CmdArgument, factory models.JsonReaderFactory) (models.Operation, error) {
	isImages := slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
		return o.Name == data.ImagesOpt
	})
	isHelp := slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
		return o.Name == data.HelpOpt
	})
	isListImages := slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
		return o.Name == data.ListImagesOpt
	})

	if isImages {
		return imagesbased.Create(arg), nil
	} else if isHelp {
		return helpbased.Create(arg), nil
	} else if isListImages {
		h, e := factory.GetReader()
		if e != nil {
			return nil, e
		}
		return listimages.Create(h), nil
	} else {
		return folderbased.Create(arg), nil
	}
}
