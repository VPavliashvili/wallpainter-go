package listimages

import (
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
)

// should create pipereceiver interface and pipesender
// this should be sender and other commands will be receivers
// this module will create a pipe and pass informationc to it
// pipelistener objects will listen to it and act accordingly during runtime
func Create(cmds.CmdArgument) models.Operation {
    return nil
}
