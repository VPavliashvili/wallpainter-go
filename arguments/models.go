package arguments

import "fmt"

type Argument interface {
	GetName() string
	String() string
	Value() string
	Description() string
}

type argument struct {
	name  *string
	value *string
	desc  *string
}

func (arg argument) GetName() string     { return *arg.name }
func (arg argument) Value() string       { return *arg.value }
func (arg argument) Description() string { return *arg.desc }

func (arg argument) String() string {
	return fmt.Sprintf("name '%v', value '%v'", *arg.name, *arg.value)
}

type argError struct {
	name  string
	value string
}

func (err argError) Error() string {
	return fmt.Sprintf("argError {invalid name: {%v} or value: {%v}}", err.name, err.value)
}
