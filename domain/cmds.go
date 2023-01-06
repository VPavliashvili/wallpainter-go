package domain

import (
	"reflect"
	"sort"

	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

type Command interface {
	SetArgument(CmdArgument)
	Execute() error
	Name() string
}

type AvailableCommandsProvider interface {
	Get() []Command
}

type CmdArgument struct {
	Flag        flags.Flag
	Opts        []opts.Opt
	Description string
}

func (arg *CmdArgument) Equals(other *CmdArgument) bool {
	if arg == nil || other == nil {
		return false
	}

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
