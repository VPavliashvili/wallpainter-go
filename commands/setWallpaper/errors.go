package setwallpaper

import "fmt"

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
