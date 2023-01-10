package setwallpaper

const Scaling = "--scaling"
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
