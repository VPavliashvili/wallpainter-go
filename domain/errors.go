package domain

import (
	"fmt"
	"reflect"

	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

type NonExistentCommandError struct {
	Flag flags.Flag
}

func (err NonExistentCommandError) Error() string {
	return fmt.Sprintf("command with argument '%v' does not exist\ntype %v to view available commands", err.Flag, flags.Help)
}

func (err NonExistentCommandError) Is(target error) bool {
	switch target := target.(type) {
	case NonExistentCommandError:
		return err.Flag == target.Flag
	default:
		panic("only NonExistentCommandError error is expected")
	}
}

type InvalidOptionsError struct {
	OptArgs []string
    OverridenMsg string
}

func (err InvalidOptionsError) Error() string {
    if err.OverridenMsg != "" {
        return err.OverridenMsg
    }
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

type EmptyInputError struct{}

func (err EmptyInputError) Error() string {
	return fmt.Sprintf("argument input is empty\nsee %v to view possible options", flags.Help)
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
