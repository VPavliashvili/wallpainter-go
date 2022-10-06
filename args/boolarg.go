package args

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type boolval struct {
	value *bool
}

func (b boolval) String() string {
	if b.value != nil {
		return fmt.Sprint(*b.value)
	}
	return "<nil>"
}

func (b *boolval) Set(s string) error {
	val, err := strconv.ParseBool(s)

	if err != nil {
		return err
	}
	b.value = &val
	return nil
}

type boolarg struct {
	value *boolval
	names *[]string
    description string
}

func (b boolarg) String() string {
    return fmt.Sprintf("names: %v\nvalue: %v\ndesc: %v", b.names, *(b.value.value), b.description)
}

func (b boolarg) GetNames() []string {
	return *b.names
}

func (b boolarg) getValue() flag.Value {
	return b.value
}

func (b boolarg) Value() string {
    return fmt.Sprint(*(b.value.value))
}

func (b boolarg) Set(s string) error {
    res, err := strconv.ParseBool(s)

    if err != nil {
        return err
    }

    b.value.value = setValue(res)
    return nil
}

func (b boolarg) GetDescription() string {
	builder := strings.Builder{}

	for i, v := range *b.names {
		if i%2 == 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(v)
	}
	formatted := builder.String()

    return fmt.Sprintf("%v\n        %v\n", formatted, b.description)
	//return fmt.Sprintf("%v\n      print this help message", formatted)
}
