package setwallpaper

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func Create() domain.Command {
	return &setWallpaper{
		io: io{},
	}
}

type setWallpaper struct {
	io      io
	imgPath string
	scaling string
}

func (s setWallpaper) Execute() error {
	if !s.io.Exist(s.imgPath) {
		return domain.InvalidPathError{Path: s.imgPath}
	} else if !s.io.IsPicture(s.imgPath) {
		return domain.NotPictureError{File: s.imgPath}
	}

	err := s.io.SetWallpaper(s.imgPath, s.scaling)
	if err != nil {
		return err
	}

	return nil
}

func (setWallpaper) Name() string {
	return flags.SetWallpaper
}

func (s *setWallpaper) SetArgument(arg domain.Argument) {
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
