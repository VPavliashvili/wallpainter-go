package folderbased

import "time"

type producingLogic interface {
	produceRunningPictures() string
}

type producingImpl struct{}


type wallpaperLogic interface {
	run([]string) error
}

type logic struct {
	time        time.Duration
}


type pathargument struct {
	time           time.Duration
	isRecursive    bool
	setterLogic    wallpaperLogic
	producingLogic producingLogic
    path string
}
