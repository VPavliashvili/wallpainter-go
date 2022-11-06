package configdata

import "strings"

type Flag string
type Opt struct {
    Name string
    Value string
}

type CmdData struct {
	FlagName Flag
    Opts []Opt
}

type AllCommandData interface {
	GetAllCommandData() []CmdData
}

func IsOptName(arg string) bool {
    return strings.HasPrefix(arg, "-")
}
