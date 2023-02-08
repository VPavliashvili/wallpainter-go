package folderbased

import (
	"fmt"
	"time"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"github.com/VPavliashvili/wallpainter-go/iohandler"
	"golang.org/x/exp/slices"
)

var lastSetPicture string
var channel chan string

func Create(arg cmds.CmdArgument) models.Operation {
	channel = make(chan string)

	return createArgumentWithFolderPath(arg, logic{
		path:        getFolderPath(arg.Opts),
		time:        sharedbehaviour.GetTimeOpt(arg.Opts),
		isRecursive: getRecursiveOpt(arg.Opts),
	})
}

type wallpaperLogic interface {
	run() error
}

type logic struct {
	path        string
	time        time.Duration
	isRecursive bool
}

func (l logic) run() error {
	pictures, err := iohandler.GetPictures(l.path, l.isRecursive)

	if err != nil {
		return err
	}

	lastSetPicture = sharedbehaviour.TakeRandomElement(pictures, lastSetPicture)
	wallpeperSetter := iohandler.GetWallpaperSetter()

	err = wallpeperSetter.SetWallpaper(lastSetPicture, data.ImageDefaultScaling)
	if err != nil {
		return err
	}

	for i := 0; i < int(l.time.Seconds()); i++ {
		time.Sleep(time.Second)
	}

	return l.run()
}

type pathargument struct {
	time        time.Duration
	isRecursive bool
	setterLogic wallpaperLogic
}

func (p pathargument) Execute() error {
	fmt.Printf("execution of folderbased started\n")

	err := p.setterLogic.run()
	if err != nil {
		return err
	}

	fmt.Printf("execution of folderbased ended\n")
	return nil
}

func createArgumentWithFolderPath(arg cmds.CmdArgument, logic wallpaperLogic) pathargument {
	res := pathargument{}
	res.time = sharedbehaviour.GetTimeOpt(arg.Opts)
	res.isRecursive = getRecursiveOpt(arg.Opts)
	res.setterLogic = logic

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
