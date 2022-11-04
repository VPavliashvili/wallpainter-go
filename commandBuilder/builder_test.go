package commandbuilder_test

import (
	"testing"

	commandbuilder "github.com/VPavliashvili/slideshow-go/commandBuilder"
	"github.com/VPavliashvili/slideshow-go/configdata"
)

const valid1 = "valid1"
const valid2 = "valid2"

type fakeCmdData struct{}

func (f fakeCmdData) Flags() []configdata.Flag {
	return []configdata.Flag{
		valid1, valid2,
	}
}

func TestShouldContainExactlyOneFlagOnZeroIndex(t *testing.T) {
	flags := fakeCmdData{}
    builder := commandbuilder.GetBuilder(flags)

	cases := []struct {
		args    []string
		isError bool
	}{
		{
			args:    []string{valid1, "idk"},
			isError: false,
		},
		{
			args:    []string{"notFlag", valid1},
			isError: true,
		},
		{
			args:    []string{valid1, valid2},
			isError: true,
		},
		{
			args:    []string{"idk", "idk"},
			isError: true,
		},
	}

	for _, item := range cases {
		got := builder.Build(item.args)
		isErr := item.isError
		if got != nil != isErr {
			var msg string
			if isErr {
				msg = "Should have returned an error"
			} else {
				msg = "Args are valid, shouln't have returned an error"
			}
			t.Errorf("%v\ngot\n%v\ncase\n%v", msg, got, item.args)
		}
	}
}
