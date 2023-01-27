package setwallpaper

import (
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/iohandler"
)

func Create() cmds.Command {
	return &setWallpaper{
		io: iohandler.Create(),
	}
}

type setWallpaper struct {
	io      iohandler.IO
	imgPath string
	scaling string
}

func (s setWallpaper) Execute() error {
	err := s.io.SetWallpaper(s.imgPath, s.scaling)
	if err != nil {
		return err
	}

	return nil
}

func (setWallpaper) Name() string {
	return flags.SetWallpaper
}

func (s *setWallpaper) SetArgument(arg cmds.CmdArgument) {
	for _, opt := range arg.Opts {
		if opt.Name == "--scaling" {
			s.scaling = opt.Value
		} else if opt.Name == "" {
			s.imgPath = opt.Value
		}
	}
	if s.scaling == "" {
		s.scaling = "scale"
	}
}
