package args_test

import "github.com/VPavliashvili/slideshow-go/cmdconf"

const (
	flag1    = "flag1"
	flag2    = "flag2"
	opt11    = "--opt11"
	opt12    = "--opt12"
	optval11 = "optval11"
)

type fakeCmdData struct{}

func (f fakeCmdData) GetAllCommandData() []cmdconf.Data{
	return []cmdconf.Data{
		{
			FlagName: flag1,
			Opts: []cmdconf.Opt{
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
			Opts:     []cmdconf.Opt{},
		},
	}
}
