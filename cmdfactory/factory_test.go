package cmdfactory_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/cmdfactory"
	"github.com/VPavliashvili/slideshow-go/domain"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		args []string
		want fakeCommand
	}{
		{
			args: []string{"flag1"},
			want: fakeCommand{flagName: "flag1"},
		},
		{
			args: []string{"flag2"},
			want: fakeCommand{flagName: "flag2"},
		},
	}

	for _, item := range cases {
		parser := fakeParser{}
        cmdfactory.Setup(fakeAvailableCommands, parser)

		got, _ := cmdfactory.Create(item.args)

		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v\n", got, want)
		}
	}
}

func TestNonExistentCommand(t *testing.T) {
	cases := []struct {
        args []string
		want   error
	}{
		{
            args: []string{"nonexistent"},
			want: domain.NonExistentCommandError{
				Argument: *getFakeArgument("nonexistent"),
			},
		},
		{
            args: []string{"flag1"},
			want:   nil,
		},
	}

	for _, item := range cases {
		parser := fakeParser{}
        cmdfactory.Setup(fakeAvailableCommands, parser)

		_, got := cmdfactory.Create(item.args)

		want := item.want

		if !errors.Is(got, want) {
			t.Errorf("error in TestNonExistentCommand\ngot\n%v\nwant\n%v", got, want)
		}
	}

}
