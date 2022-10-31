package slideshow

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

type cmdFakeArg struct {
	name  string
	value string
}

func (f cmdFakeArg) GetName() string {
	return f.name
}
func (f cmdFakeArg) String() string {
	return f.name
}
func (f cmdFakeArg) Value() string {
	return f.value
}
func (f cmdFakeArg) Description() string {
	return "fake"
}

func TestStringCmd(t *testing.T) {
	cmd := slideshow{}
	got := cmd.String()
	want := "slideshowCommand"
	if got != want {
		t.Errorf("error in String()\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestArgNamesCmd(t *testing.T) {
	want := [][]string{
		{"--path", "-p"},
		{"-t"},
		{"-r"},
	}
	cmd := slideshow{}
	got := cmd.ArgNames()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Invalid ArgNames()\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestSetArguments(t *testing.T) {
	fake := []struct {
		args []arguments.Argument
		want []arguments.Argument
	}{
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "-p"},
				cmdFakeArg{name: "-t"},
				cmdFakeArg{name: "-r"},
			},
			want: []arguments.Argument{
				cmdFakeArg{name: "-p"},
				cmdFakeArg{name: "-t"},
				cmdFakeArg{name: "-r"},
			},
		},
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "--other"},
				cmdFakeArg{name: "--path"},
			},
			want: []arguments.Argument{
				cmdFakeArg{name: "--path"},
			},
		},
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "--path"},
				cmdFakeArg{name: "-p"},
			},
			want: []arguments.Argument{
				cmdFakeArg{name: "--path"},
			},
		},
	}

	for _, item := range fake {
		cmd := &slideshow{}
		cmd.SetArguments(item.args)
		want := item.want
		got := cmd.arguments

		if !reflect.DeepEqual(got, want) {
			t.Errorf("SetArguments() error\ngot\n%v\nwant\n%v", got, want)
		}

	}
}

func assertEqual[T string | int | bool](got T, want T, t *testing.T) {
	if got != want {
		t.Errorf("setValue() error\ngot\n%v\nwant\n%v", got, want)
	}
}

func TestRequiredArgSetting(t *testing.T) {

	fake := []struct {
		args  []arguments.Argument
		wantP string
		wantT int
		wantR bool
	}{
		{
			args: []arguments.Argument{
				cmdFakeArg{name: "-p", value: "someFilePath"},
				cmdFakeArg{name: "-t", value: "30"},
				cmdFakeArg{name: "-r", value: "true"},
			},
			wantP: "someFilePath",
			wantT: 30,
			wantR: true,
		},
	}

	for _, item := range fake {
		cmd := &slideshow{}
		cmd.SetArguments(item.args)
		assertEqual(cmd.path, item.wantP, t)
		assertEqual(cmd.time, item.wantT, t)
		assertEqual(cmd.recursive, item.wantR, t)
	}

}
