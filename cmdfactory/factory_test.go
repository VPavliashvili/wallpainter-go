package cmdfactory_test

import (
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmdfactory"
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
				opts: []opts.Opt{{
					Name:  "d",
					Value: "k",
				}},
			},
		},
		{
			args: []string{"flag2"},
			want: fakeCommand{
				flagName: "flag2",
				opts: []opts.Opt{
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
