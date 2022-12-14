package newfactory_test

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/slideshow-go/newfactory"
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

		newfactory.Setup(provider)
		got := newfactory.Create(*parsedArg)

		want := item.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Create\ngot\n%v\nwant\n%v\n", got, want)
		}
	}
}
