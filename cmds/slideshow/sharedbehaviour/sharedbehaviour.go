package sharedbehaviour

import (
	"strconv"
	"time"

	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func GetTimeOpt(options []opts.Opt) time.Duration {
	var timeoptAsString string
	contains := slices.ContainsFunc(options, func(o opts.Opt) bool {
		if o.Name == data.TimeOpt {
			timeoptAsString = o.Value
			return true
		}
		return false
	})

	var res int
	if contains {
		res, _ = strconv.Atoi(timeoptAsString)
	} else {
		res = data.TimeoptDefaultVal
	}

	return time.Duration(res)
}
