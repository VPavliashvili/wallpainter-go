package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

func TestShouldOnlyHaveOneArgument(t *testing.T) {
	cmd := &help{}
	args := []arguments.Argument{
		fakeArgument{"-h"},
		fakeArgument{},
	}
	cmd.SetArguments(args)
	got := len(cmd.args)
	want := 1
	if got != want {
		t.Errorf("should have thrown argument error. got: %v, want: %v", got, want)
	}
}

func TestArgumentShouldOnlyBeHelp(t *testing.T) {
	cases := []struct {
		args []arguments.Argument
		want arguments.Argument
	}{
		{
			args: []arguments.Argument{
				fakeArgument{name: "-r"},
				fakeArgument{name: "-h"},
				fakeArgument{name: "--path"},
			},
			want: fakeArgument{name: "-h"},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--help"},
			},
			want: fakeArgument{name: "--help"},
		},
		{
			args: []arguments.Argument{
				fakeArgument{name: "--nothelp"},
			},
			want: nil,
		},
	}

	for _, item := range cases {
		want := item.want
		cmd := help{}
		cmd.SetArguments(item.args)
		if want == nil {
			continue
		}
		got := cmd.args[0]

		if !reflect.DeepEqual(got, want) {
			t.Errorf("help should only take -h or --help argument from the provided arguments list, got: %v, want: %v", got, want)
		}
	}
}

func TestGetNamesInfo(t *testing.T) {
	fake := []struct {
		names []string
		want  string
	}{
		{
			names: []string{"--help", "-h"},
			want:  "{--help, -h}",
		},
		{
			names: []string{"-r"},
			want:  "{-r}",
		},
	}

	for _, v := range fake {
		want := v.want
		got := getNamesInfo(v.names)

		if got != want {
			t.Errorf("icorrect formatting of names info\ngot:  %v\nwant: %v", got, want)
		}
	}
}

func TestGetDescriptionsInfo(t *testing.T) {
	fake := []struct {
		desc string
		want string
	}{
		{
			desc: "some random nice command argument description",
			want: "some random nice command argument description",
		},
	}

	for _, v := range fake {
		want := v.want
		got := getDescriptionInfo(v.desc)

		if got != want {
			t.Errorf("incorrect formatting of description info\ngot:  ->%v\nwant: ->%v", got, want)
		}
	}
}

func TestHelpInfo(t *testing.T) {
	fake := []struct {
		info arguments.ArgInfoPair
		want string
	}{
		{
			info: arguments.ArgInfoPair{
				Names:       []string{"--cmd", "-c"},
				Description: "description for this cmd",
			},
			want: fmt.Sprintf("{--cmd, -c}\n%vdescription for this cmd\n", helpInfoTabSize),
		},
	}

	for _, v := range fake {
		want := v.want
		got := getArgumentHelp(v.info)

		if got != want {
			t.Errorf("incorrect formatting of whole help message\ngot:  \n%v\nwant: \n%v", got, want)
		}
	}

}
