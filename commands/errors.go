package commands

import (
	"fmt"
	"strings"
)

type invalidArgumentError struct {
	argName string
}

func (err invalidArgumentError) Error() string {
	return fmt.Sprintf("argument: %v does not exist", err.argName)
}

type duplicateArgumentError struct {
	argName   string
	duplicate string
}

func (err duplicateArgumentError) Error() string {
	return fmt.Sprintf("argname: %v is duplicate since it's same as already provided %v", err.duplicate, err.argName)
}

type emptyArgumentsError struct{}

func (err emptyArgumentsError) Error() string {
	return "0 arguments has passed, aborting operation"
}

type notImplementedError struct {
	args string
}

func (err notImplementedError) Error() string {
	return fmt.Sprintf("%v needs implementation", err.args)
}

func (err *notImplementedError) setArgs(args []string) {
    var sb strings.Builder
    for _, arg := range args {
        sb.WriteString(" " + arg)
    }
    err.args = sb.String()
}
