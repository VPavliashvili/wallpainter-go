package helpbased

import (
	"fmt"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/models"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/feh"
)

func Create(cmds.CmdArgument) models.Operation {
	return helpargument{}
}

type helpargument struct{}

func (h helpargument) Execute() error {
	msg := fmt.Sprintf(`%v command has several options
      1) run '%v %v' for this help menu

      2) run '%v /images/containing/folder/' to simply run this command for all images from the directory
            example: %v ~/Pictures/

      3) run '%v /folder/ %v' to run same command as above but recursively for all subfolders
            example: %v ~/Pictures/ %v 

      4) run '%v %v [and specify distinct image file separated by space]'
            example: %v %v ~/Pictures/image1.jpg ~/Pictures/otherpic.jpg
         also feh specific scaling options can be used {%v}, default value is %v when ignored
            example: %v %v ~/pic1.jpg %v ~/pic2.jpg %v
        
      5) '%v' option can be used for every combination of this command to specify time for next wallpaper change.
         it expects value in minutes or seconds and can be passed int with prefix 'm' or 's'. default value is %v minutes
            example: %v /folder/ %v 10m (this means change folpapaer from /folder/ in every 10 minutes`,
		data.Flag, data.Flag, data.HelpOpt, data.Flag, data.Flag, data.Flag, data.Recursiveopt, data.Flag, data.Recursiveopt,
		data.Flag, data.ImagesOpt, data.Flag, data.ImagesOpt, feh.GetOptionAsString(), feh.Scale,
		data.Flag, data.ImagesOpt, feh.Center, feh.Max, data.TimeOpt, data.TimeoptDefaultVal, data.Flag, data.TimeOpt)

	fmt.Println(msg)
	return nil
}
