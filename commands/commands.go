package commands

import(
    "github.com/vpavliashvili/slideshow-go/args"
)

type Command[T int | string | bool] struct {
	Arg args.Argument[T]
}
