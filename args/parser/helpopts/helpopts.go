package helpopts

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	options "github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func Create() options.OptParser {
    return parser{}
}

type parser struct {}

func (p parser) Parse(opts []string) ([]options.Opt, error) {
    if len(opts) != 0 {
        return nil, domain.InvalidOptionsError{
            OptArgs: opts,
        }
    }

    return []options.Opt{}, nil
}
