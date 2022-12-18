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
		parser fakeParser
		want   fakeCommand
	}{
		{
			parser: fakeParser{fakearg: "flag1"},
			want:   fakeCommand{flagName: "flag1"},
		},
		{
			parser: fakeParser{fakearg: "flag2"},
			want:   fakeCommand{flagName: "flag2"},
		},
	}

	for _, item := range cases {
		args := []string{item.parser.fakearg}
		parsedArg, _ := item.parser.Parse(args)

		cmdfactory.Setup(fakeAvailableCommands)
		got, _ := cmdfactory.Create(*parsedArg)

		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v\n", got, want)
		}
	}
}

func TestNonExistentCommand(t *testing.T) {
    cases := []struct{
        parser fakeParser
        want error
    }{
        {
            parser: fakeParser{fakearg: "nonexistent"},
            want: domain.NonExistentCommandError{
                Argument: *getFakeArgument("nonexistent"),
            },
        },
        {
            parser: fakeParser{fakearg: "flag1"},
            want: nil,
        },
    }

    for _, item := range cases{
        args := []string{item.parser.fakearg}
        parsedArg, _ := item.parser.Parse(args)

        cmdfactory.Setup(fakeAvailableCommands)
        _, got := cmdfactory.Create(*parsedArg)

        want := item.want

        if !errors.Is(got, want){
            t.Errorf("error in TestNonExistentCommand\ngot\n%v\nwant\n%v", got, want)
        }
    }

}

