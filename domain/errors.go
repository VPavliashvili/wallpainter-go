package domain

import (
	"fmt"
	"reflect"

	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

type NonExistentCommandError struct {
	Argument CmdArgument
	Flag     flags.Flag
}

func (err NonExistentCommandError) Error() string {
	if err.Flag != "" {
		return fmt.Sprintf("command with argument '%v' does not exist", err.Flag)
	}
	return fmt.Sprintf("command with argument '%v' does not exist", err.Argument)
}

func (err NonExistentCommandError) Is(target error) bool {
	switch target := target.(type) {
	case NonExistentCommandError:
		return err.Argument.Flag == target.Argument.Flag || err.Flag == target.Flag
	default:
		panic("only NonExistentCommandError error is expected")
	}
}

type InvalidOptionsError struct {
	OptArgs []string
}

func (err InvalidOptionsError) Error() string {
	return fmt.Sprintf("Options -> [%v] are invalid for this command", err.OptArgs)
}

func (err InvalidOptionsError) Is(target error) bool {
	switch target := target.(type) {
	case InvalidOptionsError:
		return reflect.DeepEqual(err.OptArgs, target.OptArgs) ||
			(len(err.OptArgs) == 0 && len(target.OptArgs) == 0)
	default:
		panic("only InvalidOptionsError is expected here")
	}
}

type MoreThanOneFlagError struct {
	Args []string
}

func (err MoreThanOneFlagError) Error() string {
	return fmt.Sprintf("There is more than one command flag in input\n%v", err.Args)
}

func (err MoreThanOneFlagError) Is(target error) bool {
	switch target := target.(type) {
	case MoreThanOneFlagError:
		return reflect.DeepEqual(err.Args, target.Args)
	default:
		panic("only MoreThanOneFlagError is expected")
	}
}

type InvalidPathError struct {
	Path string
}

func (i InvalidPathError) Error() string {
	return fmt.Sprintf("file -> %v does not exist", i.Path)
}

type NotPictureError struct {
	File string
}

func (i NotPictureError) Error() string {
	return fmt.Sprintf("File -> %v is not an image", i.File)
}
