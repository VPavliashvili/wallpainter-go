package cmdfactory_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmdfactory"
	"github.com/VPavliashvili/wallpainter-go/domain"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		args []string
		want fakeCommand
	}{
		{
			args: []string{"flag1"},
			want: fakeCommand{
				flagName: "flag1",
				opts:     []domain.Opt{{
					Name:  "d",
					Value: "k",
				}},
			},
		},
		{
			args: []string{"flag2"},
			want: fakeCommand{
				flagName: "flag2",
				opts:     []domain.Opt{
                    {
                    	Name:  "opt1",
                    	Value: "val1",
                    },
                },
			},
		},
	}

	for _, item := range cases {
		parser := fakeParser{}
		factory := cmdfactory.Create(fakeAvailableCommands, parser)

		got, _ := factory.CreateCommand(item.args)

		want := &item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v\n", got, want)
		}
	}
}

func TestNonExistentCommand(t *testing.T) {
	cases := []struct {
		args []string
		want error
	}{
		{
			args: []string{"nonexistent"},
			want: domain.NonExistentCommandError{
				Argument: *getFakeArgument("nonexistent", []domain.Opt{}),
			},
		},
		{
			args: []string{"flag1"},
			want: nil,
		},
	}

	for _, item := range cases {
		parser := fakeParser{}
		factory := cmdfactory.Create(fakeAvailableCommands, parser)

		_, got := factory.CreateCommand(item.args)

		want := item.want

		if !errors.Is(got, want) {
			t.Errorf("error in TestNonExistentCommand\ngot\n%v\nwant\n%v", got, want)
		}
	}

}
