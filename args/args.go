package args

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

type Argument interface {
	GetNames() []string
	GetDescription() string
	String() string
    Value() string
    Set(string) error

    getValue() flag.Value
}

type argsProvider interface {
    getArgs() []string
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
	var result Argument

    help := getHelp()
    path := getPath()

	if slices.Contains(*help.names, key) {
		result = help
	} else if slices.Contains(*path.names, key) {
		result = path
		result.getValue().Set(value)
	} else {
		return nil, fmt.Errorf("argument %v does not exist", key)
	}

	return result, nil
}

func setValue[T int | string | bool](t T) *T {
	return &t
}

func setNames(names []string) *[]string {
	return &names
}
