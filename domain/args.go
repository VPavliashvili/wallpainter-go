package domain

import (
	"strings"

	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)



func IsOptName(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

type RawArgsProvider interface {
    Get() []RawArgument
}

type RawArgument struct {
	Flag       string
	OptsParser opts.OptParser
}
