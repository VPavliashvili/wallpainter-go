package commandbuilder_test

import (
	"fmt"
	"testing"

	commandbuilder "github.com/VPavliashvili/slideshow-go/commandBuilder"
)

func testCase(t *testing.T, item struct {
	args    []string
	isError bool
}, builder commandbuilder.Builder) (bool, string) {
	got := builder.Build(item.args)
	isErr := item.isError

	if got != nil != isErr {
		var msg string
		if isErr {
			msg = "Should have returned an error"
		} else {
			msg = "Args are valid, shouln't have returned an error"
		}
		res := fmt.Sprintf("%v\ngot\n%v\ncase\n%v", msg, got, item.args)
		return true, res
	}
	return false, ""
}

func TestShouldReturnErrorWhenNotStartingWithFlag(t *testing.T) {
	commandsData := fakeCmdData{}
	builder := commandbuilder.GetBuilder(commandsData)

	cases := []struct {
		args    []string
		isError bool
	}{
		{
			args:    []string{flag1},
			isError: false,
		},
		{
			args:    []string{"notFlag"},
			isError: true,
		},
		{
			args:    []string{flag1, flag2},
			isError: true,
		},
		{
			args:    []string{"idk", "idk"},
			isError: true,
		},
	}

	for _, item := range cases {
		isFail, msg := testCase(t, item, builder)
		if isFail {
			t.Errorf(msg)
		}
	}
}

func TestShouldReturnErrorWhenIncompatibleOrInvalidInput(t *testing.T) {
	commandsData := fakeCmdData{}
	builder := commandbuilder.GetBuilder(commandsData)

	cases := []struct {
		args    []string
		isError bool
	}{
		{
			args:    []string{flag1, opt11, optval11, opt12},
			isError: false,
		},
		{
			args:    []string{flag1, opt11, optval11, opt12, "--idk"},
			isError: true,
		},
		{
			args:    []string{flag1, opt11, "", opt12},
			isError: false,
		},
		{
			args:    []string{flag1, opt11},
			isError: false,
		},
		{
			args:    []string{flag1, opt11, opt12},
			isError: false,
		},
		{
			args:    []string{flag2},
			isError: false,
		},
		{
			args:    []string{flag1, optval11}, // only optValue without actual opt is error
			isError: true,
		},
		{
			args:    []string{flag1, "--idk"}, //incompatible opt for flag is error
			isError: true,
		},
		{
			args:    []string{flag2, "idk"},
			isError: true,
		},
	}

	for _, item := range cases {
		isFail, msg := testCase(t, item, builder)
		if isFail {
			t.Errorf(msg)
		}
	}
}
