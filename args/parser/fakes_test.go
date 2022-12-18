package parser_test

import "github.com/VPavliashvili/slideshow-go/domain"

const (
	flag1    = "flag1"
	flag2    = "flag2"
	opt11    = "--opt11"
	opt12    = "--opt12"
	optval11 = "optval11"
)

type fakeCmdData struct{}

func (f fakeCmdData) Get() []domain.Argument{
	return []domain.Argument{
		{
			FlagName: flag1,
			Opts: []domain.Opt{
				{
					Name: opt11,
                    Value: optval11,
				},
				{
					Name: opt12,
                    Value: "",
				},
			},
		},
		{
			FlagName: flag2,
			Opts:     []domain.Opt{},
		},
	}
}
