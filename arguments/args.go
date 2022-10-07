package arguments

import (
	"golang.org/x/exp/slices"
)

func GetArguments(osArgs []string, trimmer OsArgsTrimmer) ([]Argument, error) {
	args := getArgsFromConsole(osArgs, trimmer)
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
    var desc string
	for _, arg := range arguments {
		if slices.Contains(arg.names, key) {
			if arg.validate(value) {
				isValid = true
                desc = arg.desc
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
        desc: &desc,
	}

	return result, nil
}

func setValue[T int | string | bool](t T) *T {
	return &t
}
