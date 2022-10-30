package factory

import (
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

func TestCheckForInvalidArguments(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want error
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--idk"},
			},
			want: invalidArgumentError{argName: "--idk"},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "-r"},
			},
			want: nil,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
				fakeArgument{name: "-h"},
				fakeArgument{name: "--path"},
				fakeArgument{name: "-p"},
			},
			want: nil,
		},
        {
            args: []arguments.Argument{
                fakeArgument{name: "--path"},
                fakeArgument{name: "~/somePath"},
                fakeArgument{name: "-t"},
                fakeArgument{name: "10"},
            },
            want: invalidArgumentError{argName: "~/somePath"},
        },
	}

	for _, item := range fake {
		got := checkForInvalidArguments(item.args)
		want := item.want
		if got != want {
			t.Errorf("failed to check argument validity\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestCheckForDuplicateArguments(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want error
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
				fakeArgument{name: "-h"},
			},
			want: duplicateArgumentError{
				argName:   "--help",
				duplicate: "-h",
			},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "-p"},
				fakeArgument{name: "--help"},
			},
			want: nil,
		},
	}

	for _, item := range fake {
		got := checkForDuplicateArguments(item.args)
		want := item.want
		if got != want {
			t.Errorf("duplicate argument detected\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
