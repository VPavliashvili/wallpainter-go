package main

import "github.com/VPavliashvili/wallpainter-go/processor"

func main() {

	//osArgs := os.Args[1:]

	//args, err := arguments.GetArguments(osArgs)
	//if err != nil {
		//panic(err)
	//}

	//path := args[0].Value()
	//value := args[1].Value()
	//recursive, _ := strconv.ParseBool(value)
	//pictures, _ := iohandler.GetPictures(path, recursive)

	//fmt.Println(pictures)

    processor.Process()
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
