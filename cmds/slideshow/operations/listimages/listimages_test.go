package listimages

import (
	"bytes"
	"errors"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
)

type fakeSlideshowInfo struct {
	isRunnigMock bool
	pictures     []string
}

func (io fakeSlideshowInfo) IsRunning() bool {
	return io.isRunnigMock
}

func (io fakeSlideshowInfo) GetSlideshowPictures() []string {
	return io.pictures
}

func TestWhenSlideshowIsNotRunning(t *testing.T) {
	want := domain.NotRunningError{OperationName: data.Flag}

	sut := create(&bytes.Buffer{}, fakeSlideshowInfo{isRunnigMock: false})
	err := sut.Execute()

	if !errors.Is(err, want) {
		t.Errorf("error in listimages create\ngot\n%v\nwant\n%v", err, want)
	}
}

func TestWhenSlideshowIsRunning(t *testing.T) {
	cases := []struct {
		want string
		info fakeSlideshowInfo
	}{
		{
			want: "",
			info: fakeSlideshowInfo{
				isRunnigMock: true,
				pictures:     []string{},
			},
		},
		{
			want: "someimage.png",
			info: fakeSlideshowInfo{
				isRunnigMock: true,
				pictures:     []string{"someimage.png"},
			},
		},
		{
			want: "image1\nimage2",
			info: fakeSlideshowInfo{
				isRunnigMock: true,
				pictures:     []string{"image1\nimage2"},
			},
		},
        {
        	want: "",
        	info: fakeSlideshowInfo{
                isRunnigMock: true,
                pictures: nil,
            },
        },
	}

	for _, item := range cases {
		output := bytes.Buffer{}
		sut := create(&output, item.info)

		err := sut.Execute()
		if err != nil {
			t.Errorf("error should have been nil\ngot -> %v", err)
		}

		got := output.String()
		want := item.want

		if got != want {
			t.Errorf("error in Execute\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
