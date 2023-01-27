package folderbased

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"github.com/VPavliashvili/wallpainter-go/iohandler"
	"golang.org/x/exp/slices"
)

func Create(arg cmds.CmdArgument) models.Operation {
	return createArgumentWithFolderPath(arg)
}

type pathargument struct {
	folderpath  string
	time        time.Duration
	isRecursive bool
}

func (p pathargument) Execute() error {

	if info, err := os.Stat(p.folderpath); os.IsNotExist(err) || !info.IsDir() {
		return domain.InvalidPathError{Path: p.folderpath}
	}

	fmt.Printf("execution of folderbased started\n")


	wallpeperSetter := iohandler.GetWallpaperSetter()
	pictures, err := iohandler.GetPictures(p.folderpath, p.isRecursive)
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	if err != nil {
		return err
	}

	index := random.Intn(len(pictures))
	pic := pictures[index]

	err = wallpeperSetter.SetWallpaper(pic, data.ImageDefaultScaling)
	if err != nil {
		return err
	}

    for i := time.Second; i <= p.time; i += time.Second {
        time.Sleep(time.Second)
        fmt.Printf("%v has passed\n", i)
    }

	fmt.Printf("execution of folderbased ended\n")

	return p.Execute()
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
