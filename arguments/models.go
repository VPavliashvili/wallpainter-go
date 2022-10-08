package arguments

import "fmt"

const (
	unaryValue = "true"
)

type Argument interface {
	GetName() string
	String() string
	Value() string
	Description() string
}

type binaryArgument struct {
	name  *string
	value *string
	desc  *string
}

func (arg binaryArgument) GetName() string     { return *arg.name }
func (arg binaryArgument) Value() string       { return *arg.value }
func (arg binaryArgument) Description() string { return *arg.desc }

func (arg binaryArgument) String() string {
	return fmt.Sprintf("name '%v', value '%v', desc: %v", *arg.name, *arg.value, *arg.desc)
}

type unaryArgument struct {
	name *string
	desc *string
}

func (arg unaryArgument) GetName() string     { return *arg.name }
func (arg unaryArgument) Value() string       { return unaryValue }
func (arg unaryArgument) Description() string { return *arg.desc }

func (arg unaryArgument) String() string {
	return fmt.Sprintf("name '%v', value '%v', desc: %v", *arg.name, unaryValue, *arg.desc)
}

type argError struct {
	name  string
	value string
}

func (err argError) Error() string {
	return fmt.Sprintf("argError {invalid name: {%v} or value: {%v}}", err.name, err.value)
}

type parseError struct {
    passedArgs []string
}

func (err parseError) Error() string {
    return fmt.Sprintf("parseError {invalid arguments: {%v}}", err.passedArgs)
}
