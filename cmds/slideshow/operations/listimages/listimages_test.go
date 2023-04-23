package listimages

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
)

type mockForSuccess struct {
	pics []string
}

func (m mockForSuccess) ReadData() (models.SlideshowDataModel, error) {
	return models.SlideshowDataModel{
		IsRunning:         true,
		SlideshowPictures: m.pics,
	}, nil
}

type mockForError struct {
	isRunning bool
	err       error
}

func (m mockForError) ReadData() (models.SlideshowDataModel, error) {
	return models.SlideshowDataModel{
		IsRunning: m.isRunning,
	}, m.err
}

func TestWhenSlideshowIsRunningAndSuccessful(t *testing.T) {
	cases := []struct {
		input mockForSuccess
		want  models.SlideshowDataModel
	}{
		{
			input: mockForSuccess{
				pics: []string{"/path/to/pic1", "/path/to/pic2"},
			},
			want: models.SlideshowDataModel{
				IsRunning:         true,
				SlideshowPictures: []string{"/path/to/pic1", "/path/to/pic2"},
			},
		},
		{
			input: mockForSuccess{
				pics: []string{"otherpic"},
			},
			want: models.SlideshowDataModel{
				IsRunning:         true,
				SlideshowPictures: []string{"otherpic"},
			},
		},
	}

	for _, item := range cases {
		sut := operation{
			dataReader: item.input,
		}

		got, err := getdataFromJsonReader(sut)
		if err != nil {
			t.Errorf("error should have been nil")
		}
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in readDataFromJsonHandler function\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestWhenJsonHandlerIsNil(t *testing.T) {
	sut := operation{
		dataReader: nil,
	}

	_, err := getdataFromJsonReader(sut)
	want := models.ListImagesInjectionError{}

	if !errors.Is(err, want) {
		t.Errorf("incorrect error in readDataFromJsonHandler\ngot\n%v\nwant\n%v", err, want)
	}
}

func TestShouldHandleErrorsWhenSlideshowIsNotRunnnigOrUnsaccessful(t *testing.T) {
	cases := []struct {
		input mockForError
		want  error
	}{
		{
			input: mockForError{
				isRunning: false,
				err:       nil,
			},
			want: models.ListImagesError{Msg: "slideshow is not running"},
		},
		{
			input: mockForError{
				isRunning: true,
				err:       models.ListImagesError{},
			},
			want: models.ListImagesError{Msg: "error when reading from handler"},
		},
	}

	for _, item := range cases {
		sut := operation{
			dataReader: item.input,
		}
		_, got := getdataFromJsonReader(sut)

		want := item.want
		if !errors.Is(got, want) {
			t.Errorf("incorrect error in readDataFromJsonHandler\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
