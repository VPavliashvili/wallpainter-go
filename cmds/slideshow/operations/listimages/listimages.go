package listimages

import (
	"fmt"
	"io"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type listimages struct {
	writer io.Writer
}

func (l listimages) Execute() error {
	if l.writer == nil {
		return WriterNilError{}
	}
	fmt.Fprint(l.writer, "test")

	return domain.NotRunningError{OperationName: data.Flag}
}

// should create pipereceiver interface and pipesender
// this should be sender and other commands will be receivers
// this module will create a pipe and pass informationc to it
// pipelistener objects will listen to it and act accordingly during runtime
func Create(arg cmds.CmdArgument) models.Operation {
	return create(nil)
}

func create(w io.Writer) listimages {
	return listimages{
		writer: w,
	}
}
