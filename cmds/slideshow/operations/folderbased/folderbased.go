package folderbased

import (
	"fmt"
	"strings"
	"time"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/ipc"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"github.com/VPavliashvili/wallpainter-go/iohandler"
	"golang.org/x/exp/slices"
)

var lastSetPicture string
var runningPictures []string

func Create(arg cmds.CmdArgument) models.ProducerOperation {
	return createArgumentWithFolderPath(
		arg,
        logic{
            time: sharedbehaviour.GetTimeOpt(arg.Opts),
        },
		producingImpl{},
	)
}

func createArgumentWithFolderPath(arg cmds.CmdArgument, logic wallpaperLogic, producinglogic producingLogic) pathargument {
    res := pathargument{}
    res.time = sharedbehaviour.GetTimeOpt(arg.Opts)
    res.isRecursive = getRecursiveOpt(arg.Opts)
    res.path = getFolderPath(arg.Opts)
    res.setterLogic = logic
    res.producingLogic = producinglogic

    return res
}

func (l logic) run(pictures []string) error {
	lastSetPicture = sharedbehaviour.TakeRandomElement(pictures, lastSetPicture)
	wallpeperSetter := iohandler.GetWallpaperSetter()

	err := wallpeperSetter.SetWallpaper(lastSetPicture, data.ImageDefaultScaling)
	if err != nil {
		return err
	}

	for i := 0; i < int(l.time.Seconds()); i++ {
		time.Sleep(time.Second)
	}

	return l.run(pictures)
}

func (p pathargument) Execute() error {
	fmt.Printf("execution of folderbased started\n")

	pictures, err := iohandler.GetPictures(p.path, p.isRecursive)
	runningPictures = append(runningPictures, pictures...)
    ipc.Write()

	if err != nil {
		return err
	}

	err = p.setterLogic.run(pictures)
	if err != nil {
		return err
	}

	fmt.Printf("execution of folderbased ended\n")
	return nil
}

func (p producingImpl) produceRunningPictures() string {
    var builder strings.Builder
    for i, item := range runningPictures {
        trail := ""
        if i != len(runningPictures)-1 {
            trail = ", "
        }
        s := fmt.Sprintf("%v%v", item, trail)
        builder.WriteString(s)
    }
    return builder.String()
}

func (p pathargument) Produce() string {
    return p.producingLogic.produceRunningPictures()
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
