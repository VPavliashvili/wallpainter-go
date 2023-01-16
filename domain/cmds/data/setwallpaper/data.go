package setwallpaper

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

const ScalingOpt = "--scaling"

const Flag = flags.SetWallpaper

var Description = fmt.Sprintf(`Sets new wallpaper
      usage: %v /some/path/img.jpg
      options: --scaling {%v, %v, %v, %v, %v}
      scale is default`, flags.SetWallpaper, feh.Max, feh.Fill, feh.Center, feh.Scale, feh.Tile)
