package folderbased

import (
	"time"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create(arg cmds.CmdArgument) models.Operation {
	return createArgumentWithFolderPath(arg)
}

type pathargument struct {
	folderpath  string
	time       time.Duration
	isRecursive bool
}

func (p pathargument) Execute() error {
    //mychan := make(chan bool)

    //go func() {
        //dur := time.Duration(p.time)
        //time.Sleep(p.time * time.Second)
    //}()

	return nil
}

func createArgumentWithFolderPath(arg cmds.CmdArgument) pathargument {
	res := pathargument{}
	res.time = sharedbehaviour.GetTimeOpt(arg.Opts)
	res.isRecursive = getRecursiveOpt(arg.Opts)
	res.folderpath = getFolderPath(arg.Opts)

	return res
}

func getRecursiveOpt(options []opts.Opt) bool {
	return slices.ContainsFunc(options, func(o opts.Opt) bool {
		return o.Name == data.Recursiveopt
	})
}

func getFolderPath(options []opts.Opt) string {
	for _, item := range options {
		if item.Name == data.FolderPathOptName {
			return item.Value
		}
	}
	panic("no way there is not folderpath in the arguments, for this case of command opts")
}
