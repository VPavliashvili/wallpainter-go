package listimages

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/ipc"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
)

type listimages struct {
	receivedMessage string
}

func (l *listimages) Receive(s string) {
	l.receivedMessage = s
}

func (l listimages) Execute() error {
	if !ipc.Exists() {
		return NotRunningError{OperationName: "other producer operation"}
	}

    msg := "\n" + ipc.Read() + "\n"
    fmt.Print(msg)

	return nil
}

func Create(arg cmds.CmdArgument) models.ReceiverOperation {
	return &listimages{}
}
