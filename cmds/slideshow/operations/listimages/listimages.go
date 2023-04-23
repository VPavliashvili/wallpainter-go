package listimages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

func Create(handler models.StoredJsonDataReader) models.Operation {
	return operation{
		dataReader: handler,
	}
}

type operation struct {
	dataReader models.StoredJsonDataReader
}

func (o operation) Execute() error {

	data, err := getdataFromJsonReader(o)

	if err != nil {
		return err
	}

	sb := strings.Builder{}
	line := fmt.Sprintf("%v run\n there are %v pictures loaded into slideshow\n", slideshow.ListImagesOpt, len(data.SlideshowPictures))
	sb.WriteString(line)
	for i, pic := range data.SlideshowPictures {
		sb.WriteString(strconv.Itoa(i) + ": " + pic)
	}

	return nil
}

func getdataFromJsonReader(o operation) (models.SlideshowDataModel, error) {
	data := models.SlideshowDataModel{}

	if o.dataReader == nil {
		return data, models.ListImagesInjectionError{}
	}

	data, err := o.dataReader.ReadData()

	if err != nil {
		return data, models.ListImagesError{Msg: "error when reading from handler"}
	}

	if !data.IsRunning {
		return data, models.ListImagesError{Msg: "slideshow is not running"}
	}

	return data, nil
}
