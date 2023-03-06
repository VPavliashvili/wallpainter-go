package listimages

import (
	"fmt"
	"io"
	"os"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/ipc"
)

type listimages struct {
	writer io.Writer
	info  ipc.SlideshowInfo 
}

func (l listimages) Execute() error {

	if !l.info.IsRunning() {
		return domain.NotRunningError{OperationName: data.Flag}
	}

	if l.writer == nil {
		return WriterNilError{}
	}

	pictures := l.info.GetSlideshowPictures()
	var concatenated string
	for i, item := range pictures {
        var nl string
        if i > 0 {
            nl = "\n"
        }
		concatenated += fmt.Sprintf("%v%v", item, nl)
	}

	fmt.Fprint(l.writer, concatenated)

	return nil
}

// should create pipereceiver interface and pipesender
// this should be sender and other commands will be receivers
// this module will create a pipe and pass informationc to it
// pipelistener objects will listen to it and act accordingly during runtime
func Create(arg cmds.CmdArgument) models.Operation {
	return create(os.Stdout, ipc.Create())
}

func create(w io.Writer, info ipc.SlideshowInfo) listimages {
	return listimages{
		writer: w,
		info:   info,
	}
}
