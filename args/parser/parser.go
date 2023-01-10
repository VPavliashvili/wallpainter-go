package parser

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/cmds"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func Create(data domain.RawArgsProvider) Parser {
	return concreteParser{
		allArgumentsData: data,
	}
}

type Parser interface {
	Parse([]string) (*cmds.CmdArgument, error)
}
type concreteParser struct {
	allArgumentsData domain.RawArgsProvider
}

func (b concreteParser) Parse(args []string) (*cmds.CmdArgument, error) {
	if len(args) == 0 {
        return nil, domain.EmptyInputError{}
	}

	argsData := b.allArgumentsData.Get()
	flag := args[0]
	optArgs := args[1:]

	err := VerifyOnlyOneFlagAtTheFirstIndex(flag, optArgs, argsData)
	if err != nil {
		return nil, err
	}

	optsparser, err := getOptsParser(flag, argsData)
	if err != nil {
		return nil, err
	}

	opts, err := optsparser.Parse(optArgs)
	if err != nil {
		return nil, err
	}

	result := &cmds.CmdArgument{
		Flag:        flags.Flag(flag),
		Opts:        opts,
	}
	return result, nil
}

func VerifyOnlyOneFlagAtTheFirstIndex(flag string, opts []string, argsData []domain.RawArgument) error {
	var allflags []string
	for _, item := range argsData {
		allflags = append(allflags, item.Flag)
	}

	if !slices.Contains(allflags, flag) {
		return domain.NonExistentCommandError{
			Flag: flags.Flag(flag),
		}
	}

	for _, opt := range opts {
		if slices.Contains(allflags, opt) {
			a := []string{flag}
			a = append(a, opts...)

			return domain.MoreThanOneFlagError{
				Args: a,
			}
		}
	}

	return nil
}

func getOptsParser(flag string, argsData []domain.RawArgument) (opts.OptParser, error) {
	var optsparser opts.OptParser
	hasfound := false
	for _, item := range argsData {
		if item.Flag == flag {
			optsparser = item.OptsParser
			hasfound = true
			break
		}
	}
	if !hasfound {
		return nil, domain.NonExistentCommandError{
			Flag: flags.Flag(flag),
		}
	}

	return optsparser, nil
}
