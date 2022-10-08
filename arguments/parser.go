package arguments

import (
	"strings"
)

type OsArgsTrimmer interface {
	Trim([]string) ([]string, error)
}

func getArgsFromConsole(osArgs []string) (map[string]string, error) {
	result := make(map[string]string)
    raw :=osArgs

	if len(raw) == 1 {
		result[raw[0]] = ""
		return result, nil
	}
	if isOptArg(raw[0]) {
        return nil, parseError{passedArgs: raw}
	}

	for i, arg := range raw {
		if isOptArg(arg) {
			opt := arg
			cmd := raw[i-1]
			if !isCmdArg(cmd) {
				return nil, parseError{passedArgs: raw}
			}
			result[cmd] = opt
		} else {
			cmd := arg
			if i < len(raw)-1 {
				if isCmdArg(raw[i+1]) {
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
