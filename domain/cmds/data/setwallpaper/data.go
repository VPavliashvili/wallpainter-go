package setwallpaper

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

const ScalingOpt = "--scaling"
const (
	scale  = "scale"
	tile   = "tile"
	center = "center"
	max    = "max"
	fill   = "fill"
)

func IsOnveOfScalingOption(input string) bool {
	return input != scale && input != tile && input != center && input != max && input != fill
}

const Flag = flags.SetWallpaper
var Description = fmt.Sprintf(`Sets new wallpaper
      usage: %v /some/path/img.jpg
      options: --scaling {%v, %v, %v, %v, %v}
      scale is default`, flags.SetWallpaper, max, fill, center, scale, tile)
