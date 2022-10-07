package args

import (
	"os"
	"strconv"
)

var arguments = []struct{
    names []string
    validate func(string) bool
}{
	{
		names: []string{"-h", "--help"},
		validate: func(s string) bool {
			_, err := strconv.ParseBool(s)
			return err == nil
		},
	},
	{
		names: []string{"-p", "--path"},
		validate: func(s string) bool {
			if _, err := os.Stat(s); !os.IsNotExist(err) {
				return true
			}
			return false
		},
	},
	{
		names: []string{"-r"},
		validate: func(s string) bool {
			_, err := strconv.ParseBool(s)
			return err == nil
		},
	},
	{
		names: []string{"-t"},
		validate: func(s string) bool {
			if _, err := strconv.Atoi(s); err == nil {
				return true
			}
			return false
		},
	},
}
