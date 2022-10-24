package commands

import (
	"errors"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

func TestGetCommand(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want Command
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--invalid"},
			},
			want: &invalidArgumentCommand{input: []string{"--invalid"}},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "-h"},
			},
			want: &help{},
		},
	}

	for _, item := range fake {
		got := GetCommand(item.args)
		want := item.want

		if got.String() != want.String() {
			t.Errorf("wrong command returned\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item.args)
		}
	}
}

func TestCreateCommand(t *testing.T) {
	fake := []struct {
		args    []arguments.Argument
		command Command
		err     error
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
			},
			command: &help{args: []arguments.Argument{}},
			err:     nil,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--nonexistent"},
			},
			command: nil,
			err:     invalidArgumentError{argName: "--nonexistent"},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
				fakeArgument{name: "-h"},
			},
			command: nil,
			err: duplicateArgumentError{
				argName:   "--help",
				duplicate: "-h",
			},
		},
		{
			args:    []arguments.Argument{},
			command: nil,
			err:     emptyArgumentsError{},
		},
		{
			args: []arguments.Argument{
				fakeArgument{"--path"},
			},
			command: nil,
			err:     notImplementedError{args: " --path"},
		},
	}

	for _, item := range fake {
		got, err := createCommand(item.args)
		wantCmd := item.command
		wantErr := item.err

		if got != nil && wantCmd != nil {
			if got.String() != wantCmd.String() {
				t.Errorf("wrong cmd returned\ngot\n%v\nwant\n%v\ncase -> %v", got, wantCmd, item)
			}
		}
		if !errors.Is(err, wantErr) && err != wantErr {
			t.Errorf("wrong error retunred\ngot\n%v\nwant\n%v\ncase -> %v", err, wantErr, item)
		}

	}

}
