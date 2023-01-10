package parser_test

import (
	"errors"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestShouldSeparateFlagAndOptsCorrectly(t *testing.T) {
	cases := []struct {
		args []string
		want cmds.CmdArgument
	}{
		{
			args: []string{flag1, opt11, optval11},
			want: cmds.CmdArgument{
				Flag:        flag1,
				Opts:        []opts.Opt{{Name: opt11, Value: optval11}},
				Description: "",
			},
		},
	}

	parser := parser.Create(fakeCmdData{})
	for _, item := range cases {
		got, err := parser.Parse(item.args)
		want := item.want

		if err != nil {
			t.Errorf("error should have been nil, got\n%v", err)
		}

		if !got.Equals(&want) {
			t.Errorf("error in args parser\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestWhenErrorShouldThrow(t *testing.T) {
	cases := []struct {
		args []string
		err  error
	}{
		{
			args: []string{"nonexietentflag"},
			err: domain.NonExistentCommandError{
				Flag: "nonexietentflag",
			},
		},
		{
			args: []string{flag1, opt11, flag2},
			err: domain.MoreThanOneFlagError{
				Args: []string{flag1, opt11, flag2},
			},
		},
	}

	parser := parser.Create(fakeCmdData{})
	for _, item := range cases {
		res, err := parser.Parse(item.args)
		want := item.err

		if res != nil {
			t.Errorf("result should have been nil, got -> %v", res)
		}

		if !errors.Is(err, want) {
			t.Errorf("proper error has not thrown\ngot\n%v\nwant\n%v", err, want)
		}
	}
}
