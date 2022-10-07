package args

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

type Argument interface {
	GetName() string
	String() string
	Value() string
	Set(string) error

	getValue() flag.Value
}

type argument struct {
	name  *string
	value *string
}

func (arg argument) GetName() string { return *arg.name }
func (arg argument) String() string {
	return fmt.Sprintf("name '%v', value '%v'", *arg.name, *arg.value)
}
func (arg argument) Value() string        { return *arg.value }
func (arg argument) Set(s string) error   { *arg.value = s; return nil }
func (arg argument) getValue() flag.Value { panic("nobody cares anymore") }

type argError struct {
	name  string
	value string
}

func (err argError) Error() string {
	return fmt.Sprintf("argError {invalid name: {%v} or value: {%v}}", err.name, err.value)
}

func GetArguments() ([]Argument, error) {
	args := getArgsFromConsole(os.Args)
	var result []Argument

	for k, v := range args {
		arg, err := createArgument(k, v)
		if err != nil {
			return nil, err
		}
		result = append(result, arg)
	}

	return result, nil
}

func createArgument(key string, value string) (Argument, error) {
	isValid := false
	for _, arg := range arguments {
		if slices.Contains(arg.names, key) {
			if arg.validate(value) {
				isValid = true
			}
			break
		}
	}

	if !isValid {
		return nil, argError{
			name:  key,
			value: value,
		}
	}

	var result Argument = argument{
		name:  &key,
		value: &value,
	}

	return result, nil
}

func setValue[T int | string | bool](t T) *T {
	return &t
}
