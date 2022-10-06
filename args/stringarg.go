package args

import (
	"flag"
	"fmt"
	"strings"
)

type stringval struct {
	value *string
}

func (s stringval) String() string {
	if s.value != nil {
		return *s.value
	}
	return "<nil>"
}

func (s *stringval) Set(prm string) error {
	s.value = &prm
	return nil
}

type stringarg struct {
	value       *stringval
	names       *[]string
	description string
}

func (s stringarg) String() string {
    return fmt.Sprintf("names: %v\nvalue: %v\ndesc: %v", s.names, *(s.value.value), s.description)
}

func (s stringarg) GetNames() []string {
	return *s.names
}

func (s stringarg) getValue() flag.Value {
	return s.value
}

func (s stringarg) Value() string {
    return *(s.value.value)
}

func (arg stringarg) Set(s string) error {
    arg.value.value = setValue(s)
    return nil
}

func (s stringarg) GetDescription() string {
	builder := strings.Builder{}

	for i, v := range *s.names {
		if i%2 == 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(v)
	}
	formatted := builder.String()

	return fmt.Sprintf("%v\n        %v\n", formatted, s.description)
}
