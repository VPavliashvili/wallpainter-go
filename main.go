package main

import (
	"fmt"

	"github.com/vpavliashvili/slideshow-go/args"
	"github.com/vpavliashvili/slideshow-go/iohandler"
)

func main() { 
    pictures, err := iohandler.GetPictures(args.Path.Value, args.Recursive.Value)

    if err != nil {
        panic(err)
    }

	for _, picture := range pictures {
		fmt.Println(picture)
	}
}
