package setwallpaper_test

import (
	"testing"

	setwallpaper "github.com/VPavliashvili/wallpainter-go/cmds/setWallpaper"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func TestName(t *testing.T) {
    sut := setwallpaper.Create()

    got := sut.Name()
    want := flags.SetWallpaper

    if got != want {
		t.Errorf("Argument Name errro\ngot\n%v\nwant\n%v", got, want)
    }
}
