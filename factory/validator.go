package factory

import (
	"github.com/VPavliashvili/slideshow-go/arguments"
	"github.com/VPavliashvili/slideshow-go/utils"
)

func checkForInvalidArguments(args []arguments.Argument) error {
	argsInfo := arguments.GetAllArgumentInfo()
	var allNames []string
	var providedNames []string

    for _, info := range argsInfo {
        allNames = append(allNames, info.Names...)
    }
    for _, arg := range args {
        providedNames = append(providedNames, arg.GetName())
    }

	for _, name := range providedNames {
		if !utils.Contains(allNames, name) {
			return invalidArgumentError{argName: name}
		}
	}
	return nil
}

func checkForDuplicateArguments(args []arguments.Argument) error {
	argsInfo := arguments.GetAllArgumentInfo()
	//tu argsIfos elementis 1-ze meti name gvxvdeba providedName-shi mashin visrolot error
	for _, info := range argsInfo {
		counter := 0
		var alreadyFound string
		for _, arg := range args {
			providedName := arg.GetName()
			if utils.Contains(info.Names, providedName) {
				counter++
				if counter > 1 {
					return duplicateArgumentError{
						argName:   alreadyFound,
						duplicate: providedName,
					}
				}
				alreadyFound = providedName
			}
		}
	}

	return nil
}
