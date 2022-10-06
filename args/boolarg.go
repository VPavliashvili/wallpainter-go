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
    return string("<nil>")
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
}

func (b boolarg) GetNames() []string {
    return *b.names
}

func (b boolarg) GetValue() flag.Value {
    return b.value
}

func (b boolarg) GetDescription() string {
    builder := strings.Builder{}

    for i, v := range *b.names {
        if i % 2 == 1 {
            builder.WriteByte('/')
        }
        builder.WriteString(v)
    }
    formatted := builder.String()

    return fmt.Sprintf("{%v}\n      print this help message", formatted)
}

func setValue(b bool) *bool {
    return &b
}

func setNames(names []string) *[]string {
    return &names
}
