package folderbased

import (
	"reflect"
	"testing"
	"time"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestSetArgumentWhenFolderPath(t *testing.T) {
	cases := []struct {
		arg  cmds.CmdArgument
		want pathargument
	}{
        {
            arg: cmds.CmdArgument{
                Flag: flags.RunSlideShow,
                Opts: []opts.Opt{
                    {
                        Name:  data.FolderPathOptName,
                        Value: "/path/",
                    },
                },
            },
            want: pathargument{
                folderpath:  "/path/",
                time:        data.TimeoptDefaultVal,
                isRecursive: data.RecursiveDefaultVal,
            },
        },
        {
            arg: cmds.CmdArgument{
                Flag: flags.RunSlideShow,
                Opts: []opts.Opt{
                    {
                        Name:  data.FolderPathOptName,
                        Value: "/path2/",
                    },
                    {
                        Name:  data.TimeOpt,
                        Value: "20m",
                    },
                },
            },
            want: pathargument{
                time:        time.Minute * 20,
                isRecursive: data.RecursiveDefaultVal,
                folderpath:  "/path2/",
            },
        },
        {
            arg: cmds.CmdArgument{
                Flag: flags.RunSlideShow,
                Opts: []opts.Opt{
                    {
                        Name:  data.FolderPathOptName,
                        Value: "/path/",
                    },
                    {
                        Name:  data.Recursiveopt,
                        Value: "",
                    },
                },
            },
            want: pathargument{
                folderpath:  "/path/",
                time:        data.TimeoptDefaultVal,
                isRecursive: true,
            },
        },
	}

	for _, item := range cases {
		got := createArgumentWithFolderPath(item.arg)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in createArgument\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
