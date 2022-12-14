package domain

import (
	"reflect"
	"sort"
)

type Argument struct {
	FlagName Flag
	Opts     []Opt
}

func (arg Argument) Equals(other Argument) bool {
	cp := arg.Opts
	cpother := other.Opts
	if len(cp) == 0 || len(cpother) == 0 {
		return len(cp) == len(cpother) && arg.FlagName == other.FlagName
	}

	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Name > cp[j].Name
	})
	sort.Slice(cpother, func(i, j int) bool {
		return cpother[i].Name > cpother[j].Name
	})

	arg.Opts = cp
	other.Opts = cpother

	return reflect.DeepEqual(arg, other)
}

type Flag string

type Opt struct {
	Name  string
	Value string
}

type AvailableArgumentsProvider interface {
	Get() []Argument
}
