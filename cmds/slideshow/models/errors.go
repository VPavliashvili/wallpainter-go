package models

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type JsonDeserializationError struct{
    WrongJson string
}
func (e JsonDeserializationError) Error() string {
    return fmt.Sprintf("error deserializing {%v}", e.WrongJson)
}

type ListImagesInjectionError struct{}

func (e ListImagesInjectionError) Error() string {
	return "listimages operation is not injected my jsonDataHandler implementation"
}

func (e ListImagesInjectionError) Is(target error) bool {
	switch target.(type) {
	case ListImagesInjectionError:
		return true
	default:
		return false
	}
}

type ListImagesError struct {
	Msg string
}

func (e ListImagesError) Error() string {
	msg := fmt.Sprintf("error when running %v, %v", slideshow.ListImagesOpt, e.Msg)
	return msg
}

func (e ListImagesError) Is(target error) bool {
	switch target := target.(type) {
	case ListImagesError:
		return target.Msg == e.Msg
	default:
		return false
	}
}
