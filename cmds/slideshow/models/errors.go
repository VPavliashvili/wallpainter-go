package models

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type ListImagesError struct{
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
