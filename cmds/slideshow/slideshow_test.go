package slideshow

import (
	"errors"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/domain"
)

func TestWhenOperationIsNil(t *testing.T) {
	sut := runslideshow{}
	sut.operation = nil

	want := domain.OperationNilError{}
	err := sut.Execute()

	if !errors.Is(err, want) {
		t.Errorf("wrong error from execute function\ngot\n%v\nwant\n%v", err, want)
	}

}
