package setwallpaper_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
	setwallpaper "github.com/VPavliashvili/slideshow-go/commands/setWallpaper"
)

func TestString(t *testing.T) {
	want := "somePath"

	cmd := setwallpaper.Create()
	fake := cmdFakeArg{
		name:  "--imgpath",
		value: "somePath",
	}
	cmd.SetArguments([]arguments.Argument{fake})
    got := cmd.String()
	if got != want {
		t.Errorf("String() error\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestArgNames(t *testing.T) {
	want := [][]string{
		{"--imgpath"},
	}
	got := setwallpaper.Create().ArgNames()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ArgNames() error\ngot\n%v\nwant\n%v", got, want)
	}

}

func TestSetArguments(t *testing.T) {
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
		cmd := setwallpaper.Create()
		cmd.SetArguments(item.args)
		got := cmd.ImgPath
		want := item.want

		if got != want {
			t.Errorf("SetArguments() error\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestExecute(t *testing.T) {
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
			want:    setwallpaper.InvalidPathError{Path: "invalidpath"},
		},
		{
			imgpath: "validfile",
			want:    setwallpaper.NotPictureError{File: "validfile"},
		},
	}

	for _, item := range fake {
		cmd := setwallpaper.Create()
        cmd.ImgPath = item.imgpath
		cmd.Setup(fakeio{})
		want := item.want
		got := cmd.Execute()

		if !errors.Is(got, want) {
			t.Errorf("Execute() should have returned an error\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item)
		}
	}
}
