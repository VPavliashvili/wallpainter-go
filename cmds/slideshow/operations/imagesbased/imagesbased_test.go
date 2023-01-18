package imagesbased

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestCreateArgumentWhenImages(t *testing.T) {
	cases := []struct {
		arg  cmds.CmdArgument
		want imagesargument
	}{
		{
			arg: cmds.CmdArgument{
				Flag: flags.RunSlideShow,
				Opts: []opts.Opt{
					{
						Name:  data.ImagesOpt,
						Value: "",
					},
					{
						Name:  "/path/",
						Value: feh.Scale,
					},
					{
						Name:  "/path2/",
						Value: feh.Scale,
					},
				},
			},
			want: imagesargument{
				time: data.TimeoptDefaultVal,
				images: []image{
					{
						path:    "/path/",
						scaling: feh.Scale,
					},
					{
						path:    "/path2/",
						scaling: feh.Scale,
					},
				},
			},
		},
		{
			arg: cmds.CmdArgument{
				Flag: flags.RunSlideShow,
				Opts: []opts.Opt{
					{
						Name:  data.ImagesOpt,
						Value: "",
					},
					{
						Name:  "/path/",
						Value: feh.Scale,
					},
					{
						Name:  data.TimeOpt,
						Value: "15",
					},
				},
			},
			want: imagesargument{
				time:   15,
				images: []image{{path: "/path/", scaling: feh.Scale}},
			},
		},
		{
			arg: cmds.CmdArgument{
				Flag: flags.RunSlideShow,
				Opts: []opts.Opt{
					{
						Name:  data.ImagesOpt,
						Value: "",
					},
					{
						Name:  "/path/",
						Value: feh.Max,
					},
				},
			},
			want: imagesargument{
				time:   data.TimeoptDefaultVal,
				images: []image{{path: "/path/", scaling: feh.Max}},
			},
		},
	}

	for _, item := range cases {
		got := createArgumentWithImages(item.arg)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in createArgument\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

