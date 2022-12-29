package domain

import (
	"reflect"
	"sort"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

type Argument struct {
    Flag flags.Flag
	Opts        []Opt
	Description string
}

func (arg Argument) Equals(other Argument) bool {
	opts := arg.Opts
	optsOther := other.Opts
	if len(opts) == 0 || len(optsOther) == 0 {
		return len(opts) == len(optsOther) && arg.Flag == other.Flag
	}

	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Name > opts[j].Name
	})
	sort.Slice(optsOther, func(i, j int) bool {
		return optsOther[i].Name > optsOther[j].Name
	})

	arg.Opts = opts
	other.Opts = optsOther

	return reflect.DeepEqual(arg, other)
}

type ArgName string

type Opt struct {
	Name  string
	Value string
}

type AvailableArgumentsProvider interface {
	Get() []Argument
}

func IsOptName(arg string) bool {
	return strings.HasPrefix(arg, "-")
}
