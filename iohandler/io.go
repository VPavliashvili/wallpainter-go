package iohandler

import (
	"fmt"
	"os/exec"

	"github.com/VPavliashvili/wallpainter-go/domain"
)

func GetWallpaperSetter() WallpaperSetter {
    return wallpaperSetter{}
}

type wallpaperSetter struct{}

func (i wallpaperSetter) exist(file string) bool {
    return Exists(file)
}

func (i wallpaperSetter) SetWallpaper(file string, scaling string) error {
	if !i.exist(file) {
		return domain.InvalidPathError{Path: file}
	} else if !isPicture(file) {
		return domain.NotPictureError{File: file}
	}

	scaling = fmt.Sprintf("--bg-%v", scaling)
	args := []string{"--no-fehbg", scaling, file}
    fmt.Println(args)
	arr, err := exec.Command("feh", args...).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(arr[:]))
	return nil
}
