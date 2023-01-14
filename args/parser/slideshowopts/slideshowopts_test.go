package slideshowopts_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/args/parser/slideshowopts"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestWhenParsingOnlyFolderOpt(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"/path/to/some/folder/"},
			want: []opts.Opt{
				{
					Name:  "",
					Value: "/path/to/some/folder/",
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

func TestWhenParsingRecursive(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"/path/to/some/folder/", "-r"},
			want: []opts.Opt{
				{
					Name:  "",
					Value: "/path/to/some/folder/",
				},
				{
					Name:  "",
					Value: "-r",
				},
			},
		},
		{
			opts: []string{"-r", "/path/to/some/folder/"},
			want: []opts.Opt{
				{
					Name:  "",
					Value: "-r",
				},
				{
					Name:  "",
					Value: "/path/to/some/folder/",
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

func TestWhenParsingWithImageOpt(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"--images", "/path/"},
			want: []opts.Opt{
				{
					Name:  "--images",
					Value: "/path/",
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

func TestWhenOnlyOneArgButError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
		{ //this checks not for certainly if folder exists but if string looks like folder path
			opts: []string{"not a folder path"},
			err: domain.InvalidPathError{
				Path: "not a folder path",
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

func TestWhenTwoArgsButError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
		{
			opts: []string{"/valid/folder", "--invalidopt"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"/valid/folder", "--invalidopt"},
			},
		},
		{
			opts: []string{"--invalidopt", "/valid/folder"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"--invalidopt", "/valid/folder"},
			},
		},
		{
			opts: []string{"invalid path", slideshow.Recursiveopt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"invalid path", slideshow.Recursiveopt},
			},
		},
		{
			opts: []string{slideshow.Recursiveopt, "invalid path"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{slideshow.Recursiveopt, "invalid path"},
			},
		},
		{
			opts: []string{"invalid path", "--invalidopt"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"invalid path", "--invalidopt"},
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
