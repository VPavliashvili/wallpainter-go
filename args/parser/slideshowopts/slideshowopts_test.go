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

func TestWhenParsingOnlyFolderOpt(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"/path/to/some/folder/"},
			want: []opts.Opt{
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/some/folder/",
				},
			},
		},
		{
			opts: []string{"/path/to/folder/", data.TimeOpt, "10m"},
			want: []opts.Opt{
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/folder/",
				},
				{
					Name:  data.TimeOpt,
					Value: "10m",
				},
			},
		},
		{
			opts: []string{"/path/to/folder/", data.TimeOpt, "10s"},
			want: []opts.Opt{
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/folder/",
				},
				{
					Name:  data.TimeOpt,
					Value: "10s",
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
			opts: []string{"/path/to/some/folder/", data.Recursiveopt},
			want: []opts.Opt{
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/some/folder/",
				},
				{
					Name:  data.Recursiveopt,
					Value: "",
				},
			},
		},
		{
			opts: []string{data.Recursiveopt, "/path/to/some/folder/"},
			want: []opts.Opt{
				{
					Name:  data.Recursiveopt,
					Value: "",
				},
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/some/folder/",
				},
			},
		},
		{
			opts: []string{"/path/to/folder/", data.Recursiveopt, data.TimeOpt, "10m"},
			want: []opts.Opt{
				{
					Name:  data.FolderPathOptName,
					Value: "/path/to/folder/",
				},
				{
					Name: data.Recursiveopt,
                    Value:  "",
				},
				{
					Name:  data.TimeOpt,
					Value: "10m",
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
			opts: []string{data.ImagesOpt, "/path/"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path/",
					Value: "",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path1/", "/path2/"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path1/",
					Value: "",
				},
				{
					Name:  "/path2/",
					Value: "",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path1/", "/path2/", "/path3/"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path1/",
					Value: "",
				},
				{
					Name:  "/path2/",
					Value: "",
				},
				{
					Name:  "/path3/",
					Value: "",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path/", "max"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path/",
					Value: "max",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path1/", "fill", "/path2/"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path1/",
					Value: "fill",
				},
				{
					Name:  "/path2/",
					Value: "",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path/", "/path/"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path/",
					Value: "",
				},
			},
		},
		{
			opts: []string{data.ImagesOpt, "/path/", data.TimeOpt, "10m"},
			want: []opts.Opt{
				{
					Name:  data.ImagesOpt,
					Value: "",
				},
				{
					Name:  "/path/",
					Value: "",
				},
				{
					Name:  data.TimeOpt,
					Value: "10m",
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

func TestWhenHelpArg(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{"-h"},
			want: []opts.Opt{
				{
					Name:  data.HelpOpt,
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

func TestWhenImagesArgButError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
        {
            opts: []string{},
            err:  domain.InvalidOptionsError{},
        },
        {
            opts: []string{data.ImagesOpt},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt},
            },
        },
        {
            opts: []string{data.TimeOpt, "10m"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.TimeOpt, "10m"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "/path/", data.TimeOpt, "notint"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "/path/", data.TimeOpt, "notint"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "/path/", data.TimeOpt, "10m", data.TimeOpt, "5m"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "/path/", data.TimeOpt, "10m", data.TimeOpt, "5m"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "/path/", data.TimeOpt},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "/path/", data.TimeOpt},
            },
        },
		{
			opts: []string{data.ImagesOpt, "/path/", data.TimeOpt, "10"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{data.ImagesOpt, "/path/", data.TimeOpt, "10"},
			},
		},
        {
            opts: []string{data.ImagesOpt, "not an image path"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "not an image path"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "not an image path", "/validpath/"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "not an image path", "/validpath/"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "/path/", "but not valid feh scaling option"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "/path/", "but not valid feh scaling option"},
            },
        },
        {
            opts: []string{data.ImagesOpt, "scale", "/path/"},
            err: domain.InvalidOptionsError{
                OptArgs: []string{data.ImagesOpt, "scale", "/path/"},
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
			opts: []string{"invalid path", data.Recursiveopt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"invalid path", data.Recursiveopt},
			},
		},
		{
			opts: []string{data.Recursiveopt, "invalid path"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{data.Recursiveopt, "invalid path"},
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

func TestWhenHelpOptButError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
		{
			opts: []string{data.HelpOpt, "anyotheropt"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{data.HelpOpt, "anyotheropt"},
			},
		},
		{
			opts: []string{"anyotherButOnIndexZero", data.HelpOpt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"anyotherButOnIndexZero", data.HelpOpt},
			},
		},
		{
			opts: []string{data.ImagesOpt, data.HelpOpt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{data.ImagesOpt, data.HelpOpt},
			},
		},
		{
            opts: []string{"/path/", data.HelpOpt},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"/path/", data.HelpOpt},
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
