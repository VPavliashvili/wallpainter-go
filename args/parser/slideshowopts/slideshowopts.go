package slideshowopts

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func Create() opts.OptParser {
	return parser{}
}

type parser struct{}

func (p parser) Parse(options []string) ([]opts.Opt, error) {
	res := []opts.Opt{}

	if len(options) == 2 {
		if isRecursiveSlideShow(options) {
			res = append(res, opts.Opt{
				Name:  "",
				Value: options[0],
			})
			res = append(res, opts.Opt{
				Name:  "",
				Value: options[1],
			})
			return res, nil
		} else {
			return nil, domain.InvalidOptionsError{
				OptArgs: options,
			}
		}
	}

	folder := options[0]
	if !looksLikeAFolder(folder) {
		return nil, domain.InvalidPathError{
			Path: folder,
		}
	}

	res = append(res, opts.Opt{
		Name:  "",
		Value: folder,
	})

	return res, nil
}

func isRecursiveSlideShow(options []string) bool {
	return (looksLikeAFolder(options[0]) && options[1] == data.Recursiveopt) ||
		(looksLikeAFolder(options[1]) && options[0] == data.Recursiveopt)
}

func looksLikeAFolder(path string) bool {
	return path[0] == '~' || path[0] == '/'
}
