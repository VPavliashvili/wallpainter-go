package main

import (
	"fmt"
	"os"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

type defaultTrimmer struct{}

func (defaultTrimmer) Trim(args []string) ([]string, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("args parameter is problematic. args: %v", args)
	}
	return args[1:], nil
}

func main() {

	osArgs := os.Args
    args, err := arguments.GetArguments(osArgs, defaultTrimmer{})
    if err != nil {
        panic(err)
    }
    fmt.Println(args)

}

//func printPicturesFromDirectory() {
//pictures, err := iohandler.GetPictures(args.Path.Value, args.Recursive.Value)

//if err != nil {
//panic(err)
//}

//for _, picture := range pictures {
//fmt.Println(picture)
//}
//}
