package args

import (
	"os"
	"strconv"
)

var arguments = []struct {
	names    []string
	desc     string
	validate func(string) bool
}{
	{
		names: []string{"-h", "--help"},
		desc:  "prints this menu",
		validate: func(s string) bool {
			_, err := strconv.ParseBool(s)
			return err == nil
		},
	},
	{
		names: []string{"-p", "--path"},
        desc: "specifing folder destination which contains target wallpapers",
		validate: func(s string) bool {
			if _, err := os.Stat(s); !os.IsNotExist(err) {
				return true
			}
			return false
		},
	},
	{
		names: []string{"-r"},
        desc: "search images recursively in passed folder (by default its disabled)",
		validate: func(s string) bool {
			_, err := strconv.ParseBool(s)
			return err == nil
		},
	},
	{
		names: []string{"-t"},
        desc: "specifying update interval between slideshow images",
		validate: func(s string) bool {
			if _, err := strconv.Atoi(s); err == nil {
				return true
			}
			return false
		},
	},
}
