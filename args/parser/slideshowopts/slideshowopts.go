package slideshowopts

import (
	options "github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func Create() options.OptParser {
	return parser{}
}

type parser struct {}

func (p parser) Parse(opts []string) ([]options.Opt, error) {
    res := []options.Opt{}

    folder := opts[0]

    res = append(res, options.Opt{
    	Name:  "",
    	Value: folder,
    })

	return res, nil
}
