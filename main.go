package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//pictures, err := iohandler.GetPictures(args.Path.Value, args.Recursive.Value)

	//if err != nil {
	//panic(err)
	//}

	//for _, picture := range pictures {
	//fmt.Println(picture)
	//}

	//fmt.Println(getPassedArgsFromConsole())

	//args := os.Args
	//args = args[1:]
	//fmt.Println(args)

}

func getPassedArgsFromConsole() []string {
	var filtered []string
	for _, v := range os.Args {
		if strings.HasPrefix(v, "-") {
			filtered = append(filtered, parseArg(v))
		}
	}

	return filtered
}

func parseArg(arg string) string {
	if !strings.Contains(arg, "-") {
		panic(fmt.Sprintf("arguments should be started with '-' symbol, user typed %v instead\n", arg))
	}
	if len(arg) <= 1 {
		panic(fmt.Sprintf("argument name should be at least 2 characters in length, user typed %v instead\n", arg))
	}

	builder := strings.Builder{}
	for i := 0; i < len(arg); i++ {
		if arg[i] != '-' {
			builder.WriteByte(arg[i])
		}
	}

	return builder.String()
}
