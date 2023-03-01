package listimages

import (
	"bytes"
	"errors"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type writerStub struct{}
func (i writerStub) Write(p []byte) (int, error) { return 0, nil }

func TestWhenSlideshowIsNotRunning(t *testing.T) {
	want := domain.NotRunningError{OperationName: data.Flag}

	sut := create(&writerStub{})
	err := sut.Execute()

	if !errors.Is(err, want) {
		t.Errorf("error in listimages create\ngot\n%v\nwant\n%v", err, want)
	}
}

func TestBufferPrinting(t *testing.T) {

	output := bytes.Buffer{}
	sut := create(&output)

	sut.Execute()

	got := output.String()
	want := "test"
	if got != want {
		t.Errorf("error in buffer printing\ngot\n%v\nwant\n%v", got, want)
	}
}

