package args

import (
	"flag"
	"fmt"
	"os"
)

type Argument[T int | string | bool] struct {
	Name        string
	Description string
	Value       T
}

var Time Argument[int]
var Path Argument[string]
var Recursive Argument[bool]

func init() {
	home, _ := os.UserHomeDir()

	Time = Argument[int]{
		Name:        "t",
		Description: "wallpaper update rate in minutes",
		Value:       10,
	}
	Path = Argument[string]{
		Name:        "p",
		Description: "wallpapers folder destination",
		Value:       fmt.Sprintf("%v/Pictures/wallpapers", home),
	}
    Recursive = Argument[bool]{
        Name: "r",
        Description: "include pictures from subdirectories",
        Value: false,
    }

	flag.IntVar(&Time.Value, Time.Name, Time.Value, Time.Description)
	flag.StringVar(&Path.Value, Path.Name, Path.Value, Path.Description)
    flag.BoolVar(&Recursive.Value, Recursive.Name, Recursive.Value, Recursive.Description)
	flag.Parse()
}
