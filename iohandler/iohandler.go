package iohandler

import (
	"io/fs"
	"path/filepath"
)

var extensions = []string{
	".jpg",
	".png",
	".jpeg",
}

func isPicture(extension string) bool {
	for _, ext := range extensions {
		if ext == extension {
			return true
		}
	}
	return false
}

//this extracts all pircture files including subdirectory content
func handleRecursively(directory string) ([]string, error) {
	var pictures []string
	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && isPicture(filepath.Ext(path)) {
			pictures = append(pictures, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return pictures, nil
}

// this ignores subdirectories in the passed directory on a given function argument
func handleNonRecursively(directory string) ([]string, error) {
    rootAbs, _ := filepath.Abs(directory)
    var pictures []string
    err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        abs, _ := filepath.Abs(path)
        if d.IsDir() && abs != rootAbs {
            return fs.SkipDir
        }
        if !d.IsDir() && isPicture(filepath.Ext(path)){
            pictures = append(pictures, path)
        }
        return nil
    })
    
    if err != nil {
        return nil, err
    }

    return pictures, nil
}

func GetPictures(directory string, recursive bool) ([]string, error) {
	if recursive {
		return handleRecursively(directory)
	} else {
		return handleNonRecursively(directory)

	}
}
