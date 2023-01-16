package slideshowopts

import (
	"strconv"

	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func Create() opts.OptParser {
	return parser{}
}

type parser struct{}

func (p parser) Parse(options []string) ([]opts.Opt, error) {
	res := []opts.Opt{}

	if len(options) == 0 {
		return nil, domain.InvalidOptionsError{
			OverridenMsg: "need proper options for this command, see help",
		}
	}

	foundTimeOpt := false
	var timeopt opts.Opt
	var indexToRemove int
	for i, item := range options {

		if foundTimeOpt && item == data.TimeOpt {
			return nil, domain.InvalidOptionsError{OptArgs: options}
		} else if foundTimeOpt {
			continue
		}

		if item == data.TimeOpt {
			foundTimeOpt = true
			if len(options)-1 >= i+1 {
				next := options[i+1]

				if _, err := strconv.Atoi(next); err != nil {
					return nil, domain.InvalidOptionsError{OptArgs: options}
				}

				timeopt = opts.Opt{
					Name:  data.TimeOpt,
					Value: next,
				}

				indexToRemove = i
			} else {
				return nil, domain.InvalidOptionsError{OptArgs: options}
			}
		}
	}

	if foundTimeOpt {
		if len(options) <= 2 {
			return nil, domain.InvalidOptionsError{OptArgs: options}
		}
		options = removeIndex(options, indexToRemove+1)
		options = removeIndex(options, indexToRemove)
	}

	firstOpt := options[0]
	if firstOpt == data.ImagesOpt {
		res, err := handleImagesOpt(options)
		if foundTimeOpt {
			res = append(res, timeopt)
		}
		return res, err
	}

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

			if foundTimeOpt {
				res = append(res, timeopt)
			}
			return res, nil
		} else {
			return nil, domain.InvalidOptionsError{
				OptArgs: options,
			}
		}
	}

	if !looksLikeAFolder(firstOpt) {
		return nil, domain.InvalidPathError{
			Path: firstOpt,
		}
	}

	res = append(res, opts.Opt{
		Name:  "",
		Value: firstOpt,
	})

	if foundTimeOpt {
		res = append(res, timeopt)
	}

	return res, nil
}

func handleImagesOpt(options []string) ([]opts.Opt, error) {
	var res []opts.Opt
	err := domain.InvalidOptionsError{
		OptArgs: options,
	}

	if len(options) == 1 {
		err.OverridenMsg = "need list of image files after '--image' option"
		return nil, err
	}

	res = append(res, opts.Opt{
		Name:  data.ImagesOpt,
		Value: "",
	})
	rest := options[1:]
	for i := 0; i < len(rest); i++ {
		current := rest[i]

		if feh.IsNotOnveOfScalingOption(current) && !looksLikeAFile(current) {
			return nil, err
		}

		if !feh.IsNotOnveOfScalingOption(current) {
			if i > 0 {
				prev := rest[i-1]
				if !looksLikeAFile(prev) {
					return nil, err
				}
			} else if i == 0 {
				return nil, err
			}
		}

		if i+1 < len(rest) {
			next := rest[i+1]
			if looksLikeAFile(current) {
				if !looksLikeAFile(next) {
					if !feh.IsNotOnveOfScalingOption(next) {
						res = append(res, opts.Opt{
							Name:  current,
							Value: next,
						})
						i++
					} else {
						return nil, err
					}
				} else {
					res = append(res, opts.Opt{
						Name:  current,
						Value: "",
					})
				}
			} else {
				continue
			}
		} else {
			if looksLikeAFile(current) {
				res = append(res, opts.Opt{
					Name:  current,
					Value: "",
				})
			}
		}
	}

	res = removeDuplicates(res)

	return res, nil
}

func isRecursiveSlideShow(options []string) bool {
	return (looksLikeAFolder(options[0]) && options[1] == data.Recursiveopt) ||
		(looksLikeAFolder(options[1]) && options[0] == data.Recursiveopt)
}

func looksLikeAFolder(path string) bool {
	return path[0] == '~' || path[0] == '/'
}

func looksLikeAFile(path string) bool {
	return path[0] == '~' || path[0] == '/'
}

func removeDuplicates(target []opts.Opt) []opts.Opt {
	allkeys := make(map[opts.Opt]bool)
	list := []opts.Opt{}
	for _, item := range target {
		if _, value := allkeys[item]; !value {
			allkeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeIndex(a []string, i int) []string {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]

	return a
}
