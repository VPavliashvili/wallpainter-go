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

func (m mockForSuccess) WriteData() error {
	return nil
}

func (m mockForSuccess) ReadData() (models.SlideshowDataModel, error) {
	return models.SlideshowDataModel{
		IsRunning:       true,
		CyclingPictures: m.pics,
	}, nil
}

type mockForError struct {
	isRunning bool
	err       error
}

func (m mockForError) WriteData() error {
	return nil
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
				IsRunning:       true,
				CyclingPictures: []string{"/path/to/pic1", "/path/to/pic2"},
			},
		},
		{
			input: mockForSuccess{
				pics: []string{"otherpic"},
			},
			want: models.SlideshowDataModel{
				IsRunning:       true,
				CyclingPictures: []string{"otherpic"},
			},
		},
	}

	for _, item := range cases {
		sut := operation{
			dataHandler: item.input,
		}

		got, err := readDataFromJsonHandler(sut)
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
		dataHandler: nil,
	}

	_, err := readDataFromJsonHandler(sut)
	want := models.ListImagesError{Msg: "json handler is nil"}

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
			dataHandler: item.input,
		}
		_, got := readDataFromJsonHandler(sut)

		want := item.want
		if !errors.Is(got, want) {
			t.Errorf("incorrect error in readDataFromJsonHandler\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
