package slideshowopts_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/args/parser/slideshowopts"
	"github.com/VPavliashvili/wallpainter-go/domain"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestWhenValidArgument(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{
				data.ListImagesOpt,
			},
			want: []opts.Opt{
				{
					Name:  data.ListImagesOpt,
					Value: "",
				},
			},
		},
	}

	parser := slideshowopts.Create()
	for _, item := range cases {
		got, err := parser.Parse(item.opts)
		want := item.want

		if err != nil {
			t.Errorf("error should be nil, got -> %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in slideshow opts parsing\ngot\n%v\nwant\n%v", got, want)
		}

	}
}

func TestWhenError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
		{
			opts: []string{data.ListImagesOpt, "anyotheropt"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{data.ListImagesOpt, "anyotheropt"},
			},
		},
		{
			opts: []string{"anyotherButOnIndexZero", data.ListImagesOpt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"anyotherButOnIndexZero", data.ListImagesOpt},
			},
		},
	}

	parser := slideshowopts.Create()
	for _, item := range cases {
		res, err := parser.Parse(item.opts)
		want := item.err

		if res != nil {
			t.Errorf("result should be nil in this case, got -> %v", res)
		}
		if !errors.Is(err, want) {
			t.Errorf("should have thrown an error\ngot\n%v\nwant\n%v", err, want)
		}
	}

}
