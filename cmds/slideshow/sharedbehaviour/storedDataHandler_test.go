package sharedbehaviour_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
)

var goodJson string = `{"isRunning":true,"slideshowPictures":["pic1","pic2"]}`
var badJson string = `hardly a json`

func TestRead(t *testing.T) {
    sut := sharedbehaviour.GetDataReader(goodJson)

    want := models.SlideshowDataModel{
    	IsRunning:         true,
    	SlideshowPictures: []string{"pic1", "pic2"},
    }

    got, err := sut.ReadData()

    if err != nil {
        t.Errorf("error should have been nil")
    }

    if !reflect.DeepEqual(got, want) {
        t.Errorf("error in storedDataHandler ReadData implementation\ngot\n%v\nwant\n%v", got, want)
    }
}

func TestReadError(t *testing.T) {
    sut := sharedbehaviour.GetDataReader(badJson)
    want := models.JsonDeserializationError{
    	WrongJson: badJson,
    }
    _, err := sut.ReadData()

    if !errors.Is(want, err) {
        t.Errorf("error is not handled correctly in ReadData implementation\ngot\n%v\nwant\n%v", err, want)
    }
}
