package args

import (
	"flag"
	"fmt"

	"golang.org/x/exp/slices"
)

type Argument interface {
    GetNames() []string
    GetValue() flag.Value
    GetDescription() string
}

var help boolarg = boolarg{
    value: &boolval{value: setValue(false)},
    names: &[]string{"-h", "--help"},
}

func GetArguments() ([]Argument, error) {
    args := getArgsFromConsole()
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

    if(slices.Contains(*help.names, key)){
        result = help
    } else {
        return nil, fmt.Errorf("argument %v does not exist", key)
    }

    return result, nil
}
