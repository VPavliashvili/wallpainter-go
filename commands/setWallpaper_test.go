package commands

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

func TestStringSetWallpaper(t *testing.T) {
	want := "somePath"

	got := setWallpaper{imgPath: "somePath"}.String()
	if got != want {
		t.Errorf("String() error\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestArgNamesSetWallpaper(t *testing.T) {
	want := [][]string{
		{"--imgpath"},
	}
	got := setWallpaper{}.ArgNames()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ArgNames() error\ngot\n%v\nwant\n%v", got, want)
	}

}

func TestSetArgumentsSetWallpaper(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want string
	}{
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "--imgpath", value: "sample"},
			},
			want: "sample",
		},
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "--imgpath", value: "sample"},
				cmdFakeArg{name: "--invalid"},
			},
			want: "sample",
		},
	}

	for _, item := range fake {
		cmd := &setWallpaper{}
		cmd.SetArguments(item.args)
		got := cmd.imgPath
		want := item.want

		if got != want {
			t.Errorf("SetArguments() error\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

type fakeio struct{}

func (f fakeio) Exist(file string) bool {
	return file == "validpath.pnj" || file == "validfile"
}

func (f fakeio) IsPicture(file string) bool {
	return file == "validpath.pnj"
}
func (f fakeio) SetWallpaper(file string) error {
	return nil
}

func TestExecuteSetWallpaper(t *testing.T) {
	fake := []struct {
		imgpath string
		want    error
	}{
		{
			imgpath: "validpath.pnj",
			want:    nil,
		},
		{
			imgpath: "invalidpath",
			want:    invalidPathError{path: "invalidpath"},
		},
		{
			imgpath: "validfile",
			want:    notPictureError{img: "validfile"},
		},
	}

	for _, item := range fake {
		cmd := setWallpaper{imgPath: item.imgpath}
		cmd.setup(fakeio{})
		want := item.want
		got := cmd.Execute()

		if !errors.Is(got, want) {
			t.Errorf("Execute() should have returned an error\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item)
		}
	}
}
