package domain

import (
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)


type RawArgsProvider interface {
    Get() []RawArgument
}

type RawArgument struct {
	Flag       string
	OptsParser opts.OptParser
}
