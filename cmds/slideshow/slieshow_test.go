package slideshow

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestSetArgument(t *testing.T) {
	cases := []struct {
		arg  cmds.CmdArgument
		want argument
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
			want: argument{
				folderpath:  "/path/",
				time:        timeoptDefaultVal,
				isRecursive: recursiveDefaultVal,
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
                        Value: "20",
                    },
                },
            },
            want: argument{
                time:        20,
                isRecursive: false,
                folderpath: "/path2/",
            },
        },
	}

	for _, item := range cases {
		got := createArgument(item.arg)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in createArgument\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
