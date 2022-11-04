package commandbuilder

import (
	"fmt"
	"strings"
)

type invalidInputError struct {
	input []string
}

func (e invalidInputError) Error() string {
    return fmt.Sprintf("Invalid input: %v\nsee --help", strings.Join(e.input, " "))
}
func (e invalidInputError) Is(target error) bool {
	return target.Error() == e.Error()
}
