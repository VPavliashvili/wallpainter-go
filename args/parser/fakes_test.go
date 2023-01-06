package parser_test

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

const (
	flag1    = "flag1"
	flag2    = "flag2"
	opt11    = "--opt11"
	opt12    = "--opt12"
	optval11 = "optval11"
	set      = "--set"
	setopt   = "~/some/path/"
)

type fakeCmdData struct{}

func (f fakeCmdData) Get() []domain.RawArgument {
	return []domain.RawArgument{
		{
			Flag: flag1,
			OptsParser: fakeoptsparser{
				flag: flag1,
			},
		},
		{
			Flag: flag2,
			OptsParser: fakeoptsparser{
				flag: flag2,
			},
		},
	}
}

type fakeoptsparser struct {
	flag string
}

func (f fakeoptsparser) Parse(s []string) (res []opts.Opt, err error) {
	switch f.flag {
	case flag1:
		res = []opts.Opt{
			{
				Name:  opt11,
				Value: optval11,
			},
		}
	case flag2:
		res = []opts.Opt{
			{
				Name: opt12,
			},
		}
	}

	return res, err
}
