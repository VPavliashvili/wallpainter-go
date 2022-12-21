package domain

import "fmt"

type NonExistentCommandError struct {
	Argument Argument
}

func (err NonExistentCommandError) Error() string {
	return fmt.Sprintf("command with argument '%v' does not exist", err.Argument)
}

func (err NonExistentCommandError) Is(target error) bool {
	switch target := target.(type) {
	case NonExistentCommandError:
		return err.Argument.FlagName == target.Argument.FlagName
	default:
		panic("only NonExistentCommandError error is expected")
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
