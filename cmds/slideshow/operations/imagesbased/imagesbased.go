package imagesbased

import (
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type imagesargument struct {
	time   float64
	images []image
}

func (i imagesargument) Execute() error {
	return nil
}

type image struct {
	path    string
	scaling string
}

func Create(arg cmds.CmdArgument) models.Operation {
	return createArgumentWithImages(arg)
}

func createArgumentWithImages(arg cmds.CmdArgument) imagesargument {
	res := imagesargument{}
	res.time = sharedbehaviour.GetTimeOpt(arg.Opts)

	var images []image
	for _, item := range arg.Opts {
		if item.Name != data.ImagesOpt && item.Name != data.TimeOpt {
			var scaling string
			if item.Value == "" {
				scaling = data.ImageDefaultScaling
			} else {
				scaling = item.Value
			}
			image := image{
				path:    item.Name,
				scaling: scaling,
			}

			images = append(images, image)
		}
	}
	res.images = images

	return res
}
