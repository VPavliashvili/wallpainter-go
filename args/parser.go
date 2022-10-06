package args

import (
	"os"
	"strings"
)

func getArgsFromConsole() map[string]string {
	result := make(map[string]string)
    raw := os.Args
    raw = raw[1:]

    if len(raw) == 1 {
        result[raw[0]] = ""   
        return result
    }

	for i := 0; i < len(raw)-1; i++ {
		cur := raw[i]
		next := raw[i + 1]
		if isCommandArg(cur) {
			if !isCommandArg(next) {
				i++
				result[cur] = next
			} else {
				result[cur] = ""
			}
		} else {
			s := "current should always be commandArg, looks like iteration bug in foor loop"
			panic(s)
		}
	}

	return result


}


func isCommandArg(arg string) bool {
	return strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-")
}

// -r -p ~/Pictures -t 10
// -p ~/Pictures t -10
// -t 10 -r -p ~/Pictures
