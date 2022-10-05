package args

type Argument[T int | string | bool] struct {
	Name []string
    Value T
}

var help = Argument[bool] {
    Name: []string{"-h", "--help"},
    Value: false,
}

func GetArguments[T int | string | bool]() []Argument[T] {
	args := getArgsFromConsole()
    var result []Argument[T]

	for k, v := range args {
        result = append(result, createArgument[T](k, v))
	}
    return result
}

func createArgument[T int | string | bool](key string, value string) Argument[T] {
	var result Argument[T]

	switch key {
	case "-h", "--help":
        result = help
	}

	return result
}
