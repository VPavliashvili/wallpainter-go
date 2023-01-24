package sharedbehaviour

import (
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

	if contains {
		res, _ := data.GetDurationFromOpt(timeoptAsString)
		return res
	} else {
		return data.TimeoptDefaultVal
	}

}
