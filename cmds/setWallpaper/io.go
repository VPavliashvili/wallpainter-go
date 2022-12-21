package setwallpaper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/VPavliashvili/wallpainter-go/iohandler"
)

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

