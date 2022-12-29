package parser_test

import (
	"testing"

	"github.com/VPavliashvili/wallpainter-go/args/parser"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

func TestShouldReturnErrorWhenIncompatibleOrInvalidInput(t *testing.T) {
	commandsData := fakeCmdData{}
	parser := parser.Create(commandsData)

	cases := []struct {
		args    []string
		isError bool
	}{
		{
			args:    []string{flag1},
			isError: false,
		},
		{
			args:    []string{"notFlag"},
			isError: true,
		},
		{
			args:    []string{flag1, flag2},
			isError: true,
		},
		{
			args:    []string{"idk", "idk"},
			isError: true,
		},
		{
			args:    []string{flag1, opt11, optval11, opt12},
			isError: false,
		},
		{
			args:    []string{flag1, opt11, optval11, opt12, "--idk"},
			isError: true,
		},
		{
			args:    []string{flag1, opt11, "", opt12},
			isError: false,
		},
		{
			args:    []string{flag1, opt11},
			isError: false,
		},
		{
			args:    []string{flag1, opt11, opt12},
			isError: false,
		},
		{
			args:    []string{flag2},
			isError: false,
		},
		{
			args:    []string{flag1, optval11}, // only optValue without actual opt is error
			isError: true,
		},
		{
			args:    []string{flag1, "--idk"}, //incompatible opt for flag is error
			isError: true,
		},
		{
			args:    []string{flag2, "idk"},
			isError: true,
		},
        {
            args: []string{},
            isError: true,
        },
	}

	for _, item := range cases {
		_, got := parser.Parse(item.args)
		isErr := item.isError

		if got != nil != isErr {
			var msg string
			if isErr {
				msg = "Should have returned an error"
			} else {
				msg = "Args are valid, shouln't have returned an error"
			}
			t.Errorf("%v\ngot\n%v\ncase\n%v", msg, got, item.args)
		}
	}
}

func TestShouldReturnOsArgsAsCmdData(t *testing.T) {
	commandsData := fakeCmdData{}
	parser := parser.Create(commandsData)

	cases := []struct {
		args []string
		want *domain.Argument
	}{
		{
			args: []string{flag1, opt11, optval11, opt12},
			want: &domain.Argument{
				Flag: flags.ToFlag(flag1),
				Opts: []domain.Opt{
					{
						Name:  opt11,
						Value: optval11,
					},
					{
						Name: opt12,
					},
				},
			},
		},
		{
			args: []string{flag1, opt12},
			want: &domain.Argument{
				Flag: flags.ToFlag(flag1),
				Opts: []domain.Opt{
					{
						Name: opt12,
					},
				},
			},
		},
		{
			args: []string{flag2},
			want: &domain.Argument{
				Flag: flags.ToFlag(flag2),
				Opts:     []domain.Opt{},
			},
		},
		{
			args: []string{flag1, opt11, opt12},
			want: &domain.Argument{
				Flag: flags.ToFlag(flag1),
				Opts: []domain.Opt{
					{
						Name: opt11,
					},
					{
						Name: opt12,
					},
				},
			},
		},
		{
			args: []string{flag1, opt12, opt11, optval11},
			want: &domain.Argument{
				Flag: flags.ToFlag(flag1),
				Opts: []domain.Opt{
					{
						Name: opt12,
					},
					{
						Name:  opt11,
						Value: optval11,
					},
				},
			},
		},
		{
			args: []string{"--invalid"},
			want: nil,
		},
	}

	for _, item := range cases {
		got, _ := parser.Parse(item.args)
		want := item.want

		if want == nil {
			if got != nil {
				t.Errorf("Error in Build()\ngot\n%v\nwant\nnill\ncase\n%v", got, item.args)
			}
		} else if !got.Equals(*want) {
			t.Errorf("Error in Build()\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item.args)
		}
	}
}
