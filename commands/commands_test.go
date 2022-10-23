package commands

import (
	"reflect"
	"testing"
)

func TestArgNames(t *testing.T) {
	data := []struct {
		command Command
		want    [][]string
	}{
		{
			command: &help{},
			want: [][]string{
				{
					"-h", "--help",
				},
			},
		},
	}

	for _, item := range data {
		got := item.command.ArgNames()
		want := item.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("error in Command.ArgNames(). got: %v, want: %v", got, want)
		}
	}
}

type fakeArgument struct {
	name string
}

func (f fakeArgument) GetName() string {
	return f.name
}
func (f fakeArgument) String() string {
	return f.name
}
func (f fakeArgument) Value() string {
	return "fake"
}
func (f fakeArgument) Description() string {
	return "fake"
}
