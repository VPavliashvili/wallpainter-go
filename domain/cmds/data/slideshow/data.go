package slideshow

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

const Recursiveopt = "-r"
const ImagesOpt = "--images"
const TimeOpt = "-t"
const HelpOpt = "-h"
const ListImagesOpt = "--list-images"

const FolderPathOptName = "folderpath"

const TimeoptDefaultVal = 1 * time.Minute
const RecursiveDefaultVal = false
const ImageDefaultScaling = feh.Scale

const Second = 's'
const Minute = 'm'

const Flag = flags.RunSlideShow

var Description = fmt.Sprintf(`Runs slideshow, which changes wallpapers for every given time
      see '%v %v' for more instructions`, Flag, HelpOpt)

func GetDurationFromOpt(opt string) (time.Duration, error) {
	numpart := opt[:len(opt)-1]
	tmpart := opt[len(opt)-1]

	if tmpart != Second && tmpart != Minute {
		return 0, errors.New("")
	}

	var multiplier time.Duration
	if tmpart == Second {
		multiplier = time.Second
	} else if tmpart == Minute {
		multiplier = time.Minute
	}

	tm, err := strconv.Atoi(numpart)
	if err != nil {
		return 0, errors.New("")
	}

	return time.Duration(tm) * multiplier, nil
}
