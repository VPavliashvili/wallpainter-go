package builder_test

import (
	"fmt"
	"testing"

	"github.com/VPavliashvili/slideshow-go/commands/help/builder"
)

func TestHelpInfo(t *testing.T) {
	fake := []struct {
		names []string
		desc  string
		want  string
	}{
		{
			names: []string{"--cmd", "-c"},
			desc:  "description for this cmd",
			want:  fmt.Sprintf("{--cmd, -c}\n%vdescription for this cmd\n", builder.HelpInfoTabSize),
		},
	}

	builder := builder.Create()

	for _, v := range fake {
		want := v.want
		got := builder.GetHelp(v.names, v.desc)

		if got != want {
			t.Errorf("incorrect formatting of whole help message\ngot:  \n%v\nwant: \n%v", got, want)
		}
	}

}
