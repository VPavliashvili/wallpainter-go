package slideshow

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/domain/feh"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
)

const Recursiveopt = "-r"
const ImagesOpt = "--images"
const TimeOpt = "-t"
const HelpOpt = "-h"

const FolderPathOptName = "folderpath"

const TimeoptDefaultVal = 1.0
const RecursiveDefaultVal = false
const ImageDefaultScaling = feh.Scale

const Flag = flags.RunSlideShow
var Description = fmt.Sprintf(`Runs slideshow, which changes wallpapers for every given time
      see '%v %v' for more instructions`, Flag, HelpOpt)
