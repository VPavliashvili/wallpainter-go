package help_test

import (
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/help"
	"github.com/VPavliashvili/wallpainter-go/domain"
)

func TestArgumentName(t *testing.T) {
	help := help.Create()

	want := "--help"
	got := help.Name()

	if got != want {
		t.Errorf("Argument Name errro\ngot\n%v\nwant\n%v", got, want)
	}
}

var fakeArgument domain.Argument = domain.Argument{
	FlagName: "--fake",
	Opts:     []domain.Opt{},
}

func TestSetArgument(t *testing.T) {
    help := help.Create()

    help.SetArgument(fakeArgument)

    if !help.GetArgument().Equals(fakeArgument) {
        t.Errorf("SetArgument error\ngot\n%v\nwant\n%v", help.GetArgument(), fakeArgument)
    }
}
