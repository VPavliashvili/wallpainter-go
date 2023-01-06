package setwallpaperopts_test

import (
	"errors"
	"reflect"
	"testing"

	setwallpaperopts "github.com/VPavliashvili/wallpainter-go/args/parser/setWallpaperopts"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestParsingWhenOptsAreValid(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"imagepath"},
			want: []opts.Opt{
				{
					Name:  "",
					Value: "imagepath",
				},
			},
		},
		{
			opts: []string{"imagepath", "--scaling", "fill"},
			want: []opts.Opt{
				{
					Name:  "",
					Value: "imagepath",
				},
				{
					Name:  "--scaling",
					Value: "fill",
				},
			},
		},
		{
			opts: []string{"--scaling", "max", "imagepath"},
			want: []opts.Opt{
				{
					Name:  "--scaling",
					Value: "max",
				},
				{
					Name:  "",
					Value: "imagepath",
				},
			},
		},
	}

	parser := setwallpaperopts.Create()

	for _, item := range cases {
		got, err := parser.Parse(item.opts)
		want := item.want

		if err != nil {
			t.Errorf("error should be nil, got\n%v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in setwallpaper opts parsing\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestParsingWhenError(t *testing.T) {
	cases := []struct {
		opts []string
		err  domain.InvalidOptionsError
	}{
		{
			opts: []string{},
			err:  domain.InvalidOptionsError{},
		},
		{
			opts: []string{"only", "two"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"only", "two"},
			},
		},
		{
			opts: []string{"more", "than", "three", "opts"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"more", "than", "three", "opts"},
			},
		},
		{
			opts: []string{"three", "but", "bad"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"three", "but", "bad"},
			},
		},
		{
			opts: []string{"imagepath", "--notscaling", "whatever"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"imagepath", "--notscaling", "whatever"},
			},
		},
		{
			opts: []string{"imagepath", "--scaling", "invalid"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"imagepath", "--scaling", "invalid"},
			},
		},
	}

	parser := setwallpaperopts.Create()

	for _, item := range cases {
		res, got := parser.Parse(item.opts)
		want := item.err

		if res != nil {
			t.Errorf("result should be nil, got\n%v", res)
		}

		if !errors.Is(got, want) {
			t.Errorf("should have thrown an error\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
