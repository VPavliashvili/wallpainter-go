package builder

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/domain"
)

func Create() HelpBuilder{
    return concreteBuilder{}
}

const HelpInfoTabSize = "      "

type HelpBuilder interface {
	GetHelp(domain.Argument) string
}

type concreteBuilder struct{
}

func getNameInfo(name string) string {
	var sb strings.Builder
	sb.WriteString("{")
    sb.WriteString(name)
	sb.WriteString("}")
	return sb.String()
}

func getDescriptionInfo(desc string) string {
	return fmt.Sprintf("%v", desc)
}

func (b concreteBuilder) GetHelp(arg domain.Argument) string {
	result := fmt.Sprintf("%v\n%v%v\n", getNameInfo(string(arg.FlagName)), HelpInfoTabSize, getDescriptionInfo(arg.Description))
	return result
}
