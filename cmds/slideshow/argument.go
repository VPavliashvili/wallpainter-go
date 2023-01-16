package slideshow

import (
	"strconv"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

type argument struct {
	folderpath  string
	time        int
	isRecursive bool
}

func createArgument(arg cmds.CmdArgument) argument {
	res := argument{}

	var timeoptAsString string
	contains := slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
		if o.Name == data.TimeOpt {
			timeoptAsString = o.Value
			return true
		}
		return false
	})

	var timeopt int
	if contains {
		timeopt, _ = strconv.Atoi(timeoptAsString)
		res.time = timeopt
	} else {
		res.time = timeoptDefaultVal
	}

	for _, item := range arg.Opts {
		if item.Name == data.FolderPathOptName {
            res.folderpath = item.Value
		}
	}

	return res
}

//if !contains {
//res.Opts = append(res.Opts, opts.Opt{
//Name: data.TimeOpt,
//Value: timeoptDefaultVal,
//})
//}

//contains = slices.ContainsFunc(arg.Opts, func(o opts.Opt) bool {
//return o.Name == data.Recursiveopt
//})

//if !contains {
//res.Opts = append(res.Opts, opts.Opt{
//Name: data.Recursiveopt,
//Value: ,
//})
//}
