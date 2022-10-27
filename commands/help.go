package commands

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/slideshow-go/arguments"
)

const helpInfoTabSize = "      "

type help struct {
	args []arguments.Argument
}

func (h help) String() string {
    return "help command"
}

func (h help) ArgNames() [][]string {
	return [][]string{
		{"-h", "--help"},
	}
}

func (h *help) SetArguments(args []arguments.Argument) {
	for _, arg := range args {
		name := arg.GetName()
		if name == "-h" || name == "--help" {
			h.args = append(h.args, arg)
			return
		}
	}
}

func (h help) Execute() error {
	infos := arguments.GetAllArgumentInfo()

	var sb strings.Builder
	for _, info := range infos {
		sb.WriteString(getArgumentHelp(info))
	}

	fmt.Print(sb.String())

	return nil
}

func getNamesInfo(names []string) string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, name := range names {
		sb.WriteString(name)
		if i < len(names)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func getDescriptionInfo(desc string) string {
	return fmt.Sprintf("%v", desc)
}

func getArgumentHelp(info arguments.ArgInfoPair) string {
	result := fmt.Sprintf("%v\n%v%v\n", getNamesInfo(info.Names), helpInfoTabSize, getDescriptionInfo(info.Description))
	return result
}
