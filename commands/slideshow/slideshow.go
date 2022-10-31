package slideshow

import (
	"fmt"
	"strconv"

	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/iohandler"
	"github.com/VPavliashvili/slideshow-go/utils"
)

func Create() *slideshow{
    return &slideshow{}
}

type slideshow struct {
	arguments []arguments.Argument

	path      string
	time      int
	recursive bool
}

func (s slideshow) String() string {
	return "slideshowCommand"
}

func (s slideshow) ArgNames() [][]string {
	return [][]string{
		{"--path", "-p"},
		{"-t"},
		{"-r"},
	}
}

type argArr struct {
	args     []string
	isLocked bool
}

func (a *argArr) lock() {
	a.isLocked = true
}

func (s *slideshow) SetArguments(args []arguments.Argument) {
	var argsArr []argArr
	for _, names := range s.ArgNames() {
		a := &argArr{
			args:     names,
			isLocked: false,
		}
		argsArr = append(argsArr, *a)
	}

	for _, arg := range args {
		name := arg.GetName()

		for i := range argsArr {
			arr := &argsArr[i]
			if utils.Contains(arr.args, name) && !arr.isLocked {
				s.arguments = append(s.arguments, arg)
				arr.lock()
			}
		}
	}

    setValues(s)
}

func setValues(s *slideshow) {
	for _, arg := range s.arguments {
		name := arg.GetName()
		value := arg.Value()

		switch name {
		case "--path", "-p":
			s.path = value
		case "-t":
			s.time, _ = strconv.Atoi(value)
		case "-r":
			s.recursive, _ = strconv.ParseBool(value)
		}
	}
}

func (s slideshow) Execute() error {
    pictures , _:= iohandler.GetPictures(s.path, s.recursive)
    fmt.Println(pictures)

    return nil
}
