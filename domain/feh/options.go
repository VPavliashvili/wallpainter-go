package feh

const (
	Scale  = "scale"
	Tile   = "tile"
	Center = "center"
	Max    = "max"
	Fill   = "fill"
)

func IsNotOnveOfScalingOption(input string) bool {
	return input != Scale && input != Tile && input != Center && input != Max && input != Fill
}
