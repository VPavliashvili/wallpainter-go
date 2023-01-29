package helpbased

import (
	"fmt"
	"html/template"
	"os"

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
	relativePath := "./cmds/slideshow/operations/helpbased/text.tmpl"
	tmpl, err := template.ParseFiles(relativePath)

	if err != nil {
		return err
	}

	data := struct {
		Flag              string
		HelpOpt           string
		RecursiveOpt      string
		ImagesOpt         string
		TimeOpt           string
		Minute            string
		Second            string
		TimeOptDefaultVal string
		FehDefaultVal     string
		FehValues         string
		FehMax            string
		FehCenter         string
	}{
		Flag:              data.Flag,
		HelpOpt:           data.HelpOpt,
		RecursiveOpt:      data.Recursiveopt,
		ImagesOpt:         data.ImagesOpt,
		TimeOpt:           data.TimeOpt,
		Minute:            string(data.Minute),
		Second:            string(data.Second),
		TimeOptDefaultVal: fmt.Sprintf("%v", data.TimeoptDefaultVal),
		FehDefaultVal:     feh.Scale,
		FehValues:         feh.GetOptionAsString(),
		FehMax:            feh.Max,
		FehCenter:         feh.Center,
	}

	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		return err
	}

	return nil
}
