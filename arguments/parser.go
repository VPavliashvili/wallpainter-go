package arguments

import (
	"strings"
)

func getArgsFromConsole(osArgs []string) (map[string]string, error) {
	result := make(map[string]string)

    if len(osArgs) == 0{
        return map[string]string{}, nil
    }
	if len(osArgs) == 1 {
		result[osArgs[0]] = ""
		return result, nil
	}
	if isOptArg(osArgs[0]) {
		return nil, parseError{passedArgs: osArgs}
	}

	for i, arg := range osArgs {
		if isOptArg(arg) {
			opt := arg
			cmd := osArgs[i-1]
			if !isCmdArg(cmd) {
				return nil, parseError{passedArgs: osArgs}
			}
			result[cmd] = opt
		} else {
			cmd := arg
			if i < len(osArgs)-1 {
				if isCmdArg(osArgs[i+1]) {
					result[cmd] = ""
				}
			} else {
				result[cmd] = ""
			}
		}

	}
	return result, nil
}

func isCmdArg(arg string) bool {
	return strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-")
}

func isOptArg(arg string) bool {
	return !isCmdArg(arg)
}
