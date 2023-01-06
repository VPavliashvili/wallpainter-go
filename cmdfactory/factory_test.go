package cmdfactory_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmdfactory"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
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
				opts:     []opts.Opt{{
					Name:  "d",
					Value: "k",
				}},
			},
		},
		{
			args: []string{"flag2"},
			want: fakeCommand{
				flagName: "flag2",
				opts:     []opts.Opt{
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

func TestEmptySlicesOrArray(t *testing.T) {
    parser := fakeParser{}
    factory := cmdfactory.Create(fakeAvailableCommands, parser)

    _, err := factory.CreateCommand([]string{})
    if err == nil {
        t.Errorf("should have thrown an error")
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
				Argument: *getFakeArgument("nonexistent", []opts.Opt{}),
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
