package setwallpaperopts

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/setwallpaper"
	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create() opts.OptParser {
	return parser{}
}

type parser struct{}

func (p parser) Parse(opts []string) ([]opts.Opt, error) {

	err := validateIncomingInput(opts)
	if err != nil {
		return nil, err
	}

	res := getSetWallpaperCommandOpts(opts)

	return res, nil
}

func validateIncomingInput(opts []string) error {
	err := domain.InvalidOptionsError{
		OptArgs: opts,
	}

	optscount := len(opts)

	if optscount != 3 && optscount != 1 {
		return err
	} else if optscount == 3 {
		// opt at index 0 or 1 should be equal to --scaling
		// followed with one of its options
		if !IsOptName(opts[0]) && !IsOptName(opts[1]) {
			return err
		}
		scalingopt := ""
		for _, opt := range opts {
			if opt == data.ScalingOpt {
				scalingopt = data.ScalingOpt
				break
			}
		}
		if scalingopt != data.ScalingOpt {
			for _, opt := range opts {
				if IsOptName(opt) {
					scalingopt = opt
					break
				}
			}
			err.OverridenMsg = fmt.Sprintf("%v option is not valid for this command\nsee %v", scalingopt, flags.Help)
			return err
		}

		scalingval := ""
		for i, opt := range opts {
			if opt == data.ScalingOpt {
				scalingval = opts[i+1]
				break
			}
		}
		if feh.IsNotOnveOfScalingOption(scalingval) {
			err.OverridenMsg = fmt.Sprintf("'%v' is not proper keyword for option %v", scalingval, scalingopt)
			return err
		}
	}
	return nil
}

func getSetWallpaperCommandOpts(optArgs []string) []opts.Opt {
	var result []opts.Opt
	var usedIndexes []int
	for i := 0; i < len(optArgs); i++ {
		arg := optArgs[i]
		if IsOptName(arg) {
			next := opts.Opt{Name: arg, Value: optArgs[i+1]}
			usedIndexes = append(usedIndexes, i+1)
			result = append(result, next)
		} else {
			if !slices.Contains(usedIndexes, i) {
				next := opts.Opt{Name: "", Value: arg}
				result = append(result, next)
			}
		}
	}

	return result
}

func IsOptName(arg string) bool {
	return strings.HasPrefix(arg, "-")
}
