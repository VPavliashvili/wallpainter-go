package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/iohandler"
)

type invalidPathError struct {
	path string
}

func (i invalidPathError) Error() string {
	return fmt.Sprintf("file -> %v does not exist", i.path)
}

type notPictureError struct {
	img string
}

func (i notPictureError) Error() string {
	return fmt.Sprintf("File -> %v is not an image", i.img)
}

type setWallpaper struct {
	imgPath string
	io      iohandler.IO
}

func (s *setWallpaper) setup(io iohandler.IO) {
	s.io = io
}

func (s setWallpaper) String() string {
	return s.imgPath
}

func (s setWallpaper) ArgNames() [][]string {
	return [][]string{
		{"--imgpath"},
	}
}

func (s *setWallpaper) SetArguments(args []arguments.Argument) {
	for _, arg := range args {
		if arg.GetName() == "--imgpath" {
			s.imgPath = arg.Value()
			return
		}
	}
}

type io struct{}

func (i io) Exist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func (i io) IsPicture(file string) bool {
    return iohandler.IsPicture(file)
}

func (i io) SetWallpaper(file string) error {
    args := []string{"--bg-scale", file}
    arr, err := exec.Command("feh", args...).Output()
    if err != nil {
        return err
    }
    fmt.Println(string(arr[:]))
    return nil
}

func (s setWallpaper) Execute() error {
	if !s.io.Exist(s.imgPath) {
		return invalidPathError{path: s.imgPath}
	} else if !s.io.IsPicture(s.imgPath) {
		return notPictureError{img: s.imgPath}
	}

    err := s.io.SetWallpaper(s.imgPath)
    if err != nil {
        return err
    }

	return nil
}

func createSetWallpaper() *setWallpaper {
    cmd := &setWallpaper{}
    cmd.setup(io{})
    return cmd
}
