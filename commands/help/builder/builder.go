package builder

import (
	"fmt"
	"strings"
)

func Create() HelpBuilder{
    return concreteBuilder{}
}

const HelpInfoTabSize = "      "

type HelpBuilder interface {
	GetHelp([]string, string) string
}

type concreteBuilder struct{
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

func (b concreteBuilder) GetHelp(names []string, desc string) string {
	result := fmt.Sprintf("%v\n%v%v\n", getNamesInfo(names), HelpInfoTabSize, getDescriptionInfo(desc))
	return result
}
