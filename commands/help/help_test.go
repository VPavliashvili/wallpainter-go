package help_test

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/commands/help"
)

func TestArgNames(t *testing.T) {
	want := [][]string{
		{"-h", "--help"},
	}
	got := help.Create().ArgNames()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ArgNames() error\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestString(t *testing.T) {
	want := "help command"
	got := help.Create().String()

	if want != got {
		t.Errorf("String() error\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestSetArguments(t *testing.T) {
	cases := []struct {
		args []arguments.Argument
		want bool
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "-r"},
				fakeArgument{name: "-h"},
				fakeArgument{name: "--path"},
			},
			want: false,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
				fakeArgument{name: "-h"},
			},
			want: false,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--nothelp"},
			},
			want: false,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "-h"},
			},
			want: true,
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
			},
			want: true,
		},
	}

	for _, item := range cases {
		want := item.want
		cmd := help.Create()
		cmd.SetArguments(item.args)

		got := cmd.Value

		if !reflect.DeepEqual(got, want) {
			t.Errorf("SetArguments() error\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestExecute(t *testing.T) {
	fake := []struct {
		infos []arguments.ArgInfoPair
		want  string
	}{
		{
			infos: []arguments.ArgInfoPair{
				{
					Names:       []string{"--idkName"},
					Description: "idk description",
				},
			},
			want: "name: [--idkName]\ndescription: idk description",
		},
	}

	for _, item := range fake {
		cmd := help.Create()
		cmd.Setup(fakebuilder{}, item.infos)
		want := item.want
		cmd.Execute()
		got := cmd.HelpText

		if got != want {
			t.Errorf("Execute() error\ngot\n%v\nwant\n%v\ncase\n%v", got, want, item)
		}
	}

}
