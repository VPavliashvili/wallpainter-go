package setwallpaperopts

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create() opts.OptParser {
    return parser{}
}

type parser struct {}

func (p parser) Parse(opts []string) ([]opts.Opt, error) {
	err := domain.InvalidOptionsError{
		OptArgs: opts,
	}

	optscount := len(opts)

	if optscount != 3 && optscount != 1 {
		return nil, err
	} else if optscount == 3 {
		// opt at index 0 or 1 should be equal to --scaling
		// followed with one of its options
		if !domain.IsOptName(opts[0]) && !domain.IsOptName(opts[1]) {
			return nil, err
		}
	}

	res := getSetWallpaperCommandOpts(opts)

	return res, nil
}

func getSetWallpaperCommandOpts(optArgs []string) []opts.Opt {
	var result []opts.Opt
	var usedIndexes []int
	for i := 0; i < len(optArgs); i++ {
		arg := optArgs[i]
		if domain.IsOptName(arg) {
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
