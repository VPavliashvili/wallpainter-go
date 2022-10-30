package factory

import (
	"errors"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

type fakeArgument struct {
	name string
}

func (f fakeArgument) GetName() string {
	return f.name
}
func (f fakeArgument) String() string {
	return f.name
}
func (f fakeArgument) Value() string {
	return "fake"
}
func (f fakeArgument) Description() string {
	return "fake"
}

func TestGetCommand(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want string
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--invalid"},
			},
			want: "--invalid",
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "-h"},
			},
			want: "help command",
		},
	}

	for _, item := range fake {
		got := GetCommand(item.args)
		want := item.want

		if got.String() != want {
			t.Errorf("wrong command returned\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item.args)
		}
	}
}

func TestCreateCommandValid(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want string
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
			},
			want: "help command",
		},
	}

	for _, item := range fake {
		got, _ := createCommand(item.args)
		want := item.want
		if got.String() != want {
			t.Errorf("wrong cmd returned\ngot\n%v\nwant\n%v\ncase -> %v", got, want, item)
		}
	}
}

func TestCreateCommandInvalid(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		err  error
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--nonexistent"},
			},
			err: invalidArgumentError{argName: "--nonexistent"},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
				fakeArgument{name: "-h"},
			},
			err: duplicateArgumentError{
				argName:   "--help",
				duplicate: "-h",
			},
		},
		{
			args: []arguments.Argument{},
			err:  emptyArgumentsError{},
		},
	}

	for _, item := range fake {
		got, err := createCommand(item.args)
		wantErr := item.err

		if got != nil {
			t.Errorf("command should have been nil\ngot\n%v\ncase -> %v", got, item)
		}
		if !errors.Is(err, wantErr) && err != wantErr {
			t.Errorf("wrong error retunred\ngot\n%v\nwant\n%v\ncase -> %v", err, wantErr, item)
		}

	}

}
