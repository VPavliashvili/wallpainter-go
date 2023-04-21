package slideshow_test

import (
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow"
)

func TestShouldReturnErrorWhenExecutesNullCommand(t *testing.T) {
    sut := slideshow.Create()

    err := sut.Execute()

    if err == nil {
        t.Errorf("error should have been thrown")
    }
}
