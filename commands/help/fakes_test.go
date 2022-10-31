package help_test

import "fmt"

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

type fakebuilder struct{}

func (f fakebuilder) GetHelp(names []string, desc string) string {
	return fmt.Sprintf("name: %v\ndescription: %v", names, desc)
}

