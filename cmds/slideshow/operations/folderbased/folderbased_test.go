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
				time:        data.TimeoptDefaultVal,
				isRecursive: true,
				setterLogic: nil,
			},
		},
	}

	for _, item := range cases {
		got := createArgumentWithFolderPath(item.arg, nil)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in createArgument\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
