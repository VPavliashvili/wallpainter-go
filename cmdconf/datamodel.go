package cmdconf

import (
	"reflect"
	"sort"
	"strings"
)

type Flag string
type Opt struct {
	Name  string
	Value string
}

type Data struct {
	FlagName Flag
	Opts     []Opt
}

func (cd Data) Equals(other Data) bool {
    cp := cd.Opts
    cpother := other.Opts
    if len(cp) == 0 || len(cpother) == 0 {
        return len(cp) == len(cpother) && cd.FlagName == other.FlagName
    }
    

	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Name > cp[j].Name
	})
	sort.Slice(cpother, func(i, j int) bool {
		return cpother[i].Name > cpother[j].Name
	})

	cd.Opts = cp
	other.Opts = cpother

	return reflect.DeepEqual(cd, other)
}

type AllCommandData interface {
	GetAllCommandData() []Data
}

func IsOptName(arg string) bool {
	return strings.HasPrefix(arg, "-")
}
