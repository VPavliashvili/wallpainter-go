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

type mockJsonReaderFactory struct{}

func (m mockJsonReaderFactory) GetReader() (models.StoredJsonDataReader, error) {
	return nil, nil
}

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
var outputListImages = listimages.Create(nil)

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
		}, {
			input: inputListImages,
			want:  outputListImages,
		},
	}

	mockReaderFactory := mockJsonReaderFactory{}
	for _, item := range cases {
		got, _ := operations.Create(item.input, mockReaderFactory)
		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestListImagesIsInjected(t *testing.T) {
	mockReaderFactory := mockJsonReaderFactory{}
	got, _ := operations.Create(inputImagesBased, mockReaderFactory)

	err := got.Execute()

	if _, throw := err.(*models.ListImagesInjectionError); throw {
		t.Errorf("listimages operation is not injected")
	}

}
