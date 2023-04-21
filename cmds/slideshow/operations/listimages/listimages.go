package listimages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

func Create(handler models.StoredJsonDataHandler) models.Operation {
	return operation{
		dataHandler: handler,
	}
}

type operation struct {
	dataHandler models.StoredJsonDataHandler
}

func (o operation) Execute() error {

    data, err := readDataFromJsonHandler(o)

    if err != nil {
        return err
    }

    sb := strings.Builder{}
    line := fmt.Sprintf("%v run\n there are %v pictures loaded into slideshow\n", slideshow.ListImagesOpt, len(data.CyclingPictures))
    sb.WriteString(line)
    for i, pic := range data.CyclingPictures {
        sb.WriteString(strconv.Itoa(i) + ": " + pic)
    }

	return nil
}

func readDataFromJsonHandler(o operation) (models.SlideshowDataModel, error) {
	data := models.SlideshowDataModel{}

	if o.dataHandler == nil {
		return data, models.ListImagesError{Msg: "json handler is nil"}
	}

	data, err := o.dataHandler.ReadData()

	if err != nil {
		return data, models.ListImagesError{Msg: "error when reading from handler"}
	}

    if !data.IsRunning {
        return data, models.ListImagesError{Msg: "slideshow is not running"}
    }

	return data, nil 
}
