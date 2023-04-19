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

type producingmock struct{
    value string
}
func (p producingmock) produceRunningPictures() string {
    return p.value
}

func TestProduce(t *testing.T) {
	cases := []struct {
		pm producingmock
	}{
		{
            pm: producingmock{
            	value: "hello there",
            },
		},
		{
			pm: producingmock{
				value: "test",
			},
		},
	}
	arg := cmds.CmdArgument{
        Flag: flags.RunSlideShow,
		Opts: []opts.Opt{
            {
            	Name: data.FolderPathOptName, 
            	Value: "",
            },
        },
	}

	for _, item := range cases {
		sut := createArgumentWithFolderPath(arg, nil, item.pm)
		got := sut.Produce()
		want := item.pm.value

		if got != want {
			t.Errorf("error in produce\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

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
                path: "/path/",
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
                path: "/path2/",
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
                path: "/path/",
			},
		},
	}

	for _, item := range cases {
		got := createArgumentWithFolderPath(item.arg, nil, nil)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in createArgument\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
