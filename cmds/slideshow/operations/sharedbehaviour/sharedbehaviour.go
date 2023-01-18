package sharedbehaviour

import (
	"strconv"

	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func GetTimeOpt(options []opts.Opt) float64 {
	var timeoptAsString string
	contains := slices.ContainsFunc(options, func(o opts.Opt) bool {
		if o.Name == data.TimeOpt {
			timeoptAsString = o.Value
			return true
		}
		return false
	})

	var res float64
	if contains {
		res, _ = strconv.ParseFloat(timeoptAsString, 64)
	} else {
		res = data.TimeoptDefaultVal
	}

	return res
}
