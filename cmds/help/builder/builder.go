package builder

import (
	"fmt"
	"strings"

	"github.com/VPavliashvili/wallpainter-go/domain"
)

func Create() HelpBuilder{
    return concreteBuilder{}
}

const Tab = "      "

type HelpBuilder interface {
	GetHelp(domain.CmdArgument) string
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

func (b concreteBuilder) GetHelp(arg domain.CmdArgument) string {
    flag := string(arg.Flag)
	result := fmt.Sprintf("%v\n%v%v\n", getNameInfo(flag), Tab, getDescriptionInfo(arg.Description))
	return result
}
