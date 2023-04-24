package slideshow

import (
	"errors"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func Create() cmds.Command {
	return &runslideshow{}
}

type runslideshow struct {
	operation models.Operation
    err error
}

func (r *runslideshow) SetArgument(arg cmds.CmdArgument) {
    factory := factoryimpl{}
    var e error
    r.operation, e = operations.Create(arg, factory)
    r.err = e
}

func (r runslideshow) Execute() error {
	if r.operation == nil {
		return errors.New("slideshow command encountered error")
	}
    if r.err != nil {
        return r.err
    }

	return r.operation.Execute()
}

func (r runslideshow) Name() string {
	return flags.RunSlideShow
}

type factoryimpl struct{}
func (f factoryimpl) GetReader() (models.StoredJsonDataReader, error) {
    filepath := slideshow.JsonDataFileLocation
    json, err := sharedbehaviour.GetJsonStringFromFile(filepath)
    if err != nil {
        return nil, err
    }
    handler := sharedbehaviour.GetDataReader(json)
    return handler, nil
}
