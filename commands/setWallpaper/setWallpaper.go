package setwallpaper

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/iohandler"
)

func Create() *setWallpaper{
    cmd := &setWallpaper{}
    cmd.Setup(io{})
    return cmd
}

type setWallpaper struct {
	ImgPath string
	io      iohandler.IO
}

func (s *setWallpaper) Setup(io iohandler.IO) {
	s.io = io
}

func (s setWallpaper) String() string {
    return "setWallpaper: " + s.ImgPath
}

func (s setWallpaper) ArgNames() [][]string {
	return [][]string{
		{"--imgpath"},
	}
}

func (s *setWallpaper) SetArguments(args []arguments.Argument) {
	for _, arg := range args {
		if arg.GetName() == "--imgpath" {
			s.ImgPath = arg.Value()
			return
		}
	}
}

func (s setWallpaper) Execute() error {
	if !s.io.Exist(s.ImgPath) {
		return InvalidPathError{Path: s.ImgPath}
	} else if !s.io.IsPicture(s.ImgPath) {
		return NotPictureError{File: s.ImgPath}
	}

    err := s.io.SetWallpaper(s.ImgPath)
    if err != nil {
        return err
    }

	return nil
}
