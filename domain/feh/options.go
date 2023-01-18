package feh

import "fmt"

const (
	Scale  = "scale"
	Tile   = "tile"
	Center = "center"
	Max    = "max"
	Fill   = "fill"
)

func GetOptionAsString() string {
    return fmt.Sprintf("%v, %v, %v, %v, %v", Max, Fill, Center, Scale, Tile)
}

func IsNotOnveOfScalingOption(input string) bool {
	return input != Scale && input != Tile && input != Center && input != Max && input != Fill
}
