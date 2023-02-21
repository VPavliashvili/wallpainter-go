package operations_test

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/folderbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/helpbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/imagesbased"
	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/operations/listimages"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

var inputFolderBased = cmds.CmdArgument{
	Flag: flags.RunSlideShow,
	Opts: []opts.Opt{
		{
			Name:  data.FolderPathOptName,
			Value: "/path/",
		},
	},
}
var outputFolderBased = folderbased.Create(inputFolderBased)

var inputImagesBased = cmds.CmdArgument{
	Flag: flags.RunSlideShow,
	Opts: []opts.Opt{
		{
			Name:  data.ImagesOpt,
			Value: "",
		},
		{
			Name:  "/path/",
			Value: "",
		},
	},
}
var outputImagesBased = imagesbased.Create(inputImagesBased)

var inputHelpBased = cmds.CmdArgument{
	Flag: flags.RunSlideShow,
	Opts: []opts.Opt{
        {
        	Name:  data.HelpOpt,
        	Value: "",
        },
    },
}
var outputHelpBased = helpbased.Create(inputHelpBased)

var inputListImages = cmds.CmdArgument{
	Flag: flags.RunSlideShow,
	Opts: []opts.Opt{
        {
        	Name:  data.ListImagesOpt,
        	Value: "",
        },
    },
}
var outputListImages = listimages.Create(inputHelpBased)

func TestCreate(t *testing.T) {
	cases := []struct {
		input cmds.CmdArgument
		want  models.Operation
	}{
		{
			input: inputFolderBased,
			want:  outputFolderBased,
		},
		{
			input: inputImagesBased,
			want:  outputImagesBased,
		},
        {
        	input: inputHelpBased,
        	want:  outputHelpBased,
        },{
        	input: inputListImages,
        	want:  outputListImages,
        },
	}

	for _, item := range cases {
		got := operations.Create(item.input)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
