package parser

import (
	"github.com/VPavliashvili/slideshow-go/domain"
	"golang.org/x/exp/slices"
)

func Create(data domain.AvailableArgumentsProvider) Parser {
	return concreteParser{
		allCommandsData: data,
	}
}

type Parser interface {
	Parse([]string) (*domain.Argument, error)
}

type concreteParser struct {
	allCommandsData domain.AvailableArgumentsProvider
}

func (b concreteParser) Parse(args []string) (*domain.Argument, error) {
	cmdsData := b.allCommandsData.Get()
	flag := args[0]
	optArgs := args[1:]

	if !hasOnlyOneFlagArgumentAtStartingPosition(flag, args, cmdsData) ||
		!checkAllEnteredFlagOptionArgumentsAreValid(flag, optArgs, cmdsData) {
		return nil, invalidInputError{input: args}
	}

	result := getCmddata(flag, optArgs)
	return &result, nil
}

func getCmddata(flag string, optArgs []string) domain.Argument {
	var result domain.Argument
	result.FlagName = domain.Flag(flag)
	pairs := make(map[string]string)

	for i := 0; i < len(optArgs); i++ {
		arg := optArgs[i]
		if i < len(optArgs)-1 {
			next := optArgs[i+1]
			if !domain.IsOptName(next) {
				i++
				pairs[arg] = next
			} else {
				pairs[arg] = ""
			}
		} else {
			pairs[arg] = ""
		}
	}

	for opt, val := range pairs {
		result.Opts = append(result.Opts, domain.Opt{
			Name:  opt,
			Value: val,
		})
	}

	return result
}

func checkAllEnteredFlagOptionArgumentsAreValid(flag string, optArgs []string, cmdsData []domain.Argument) bool {
	if len(optArgs) > 0 && !domain.IsOptName(optArgs[0]) {
		return false
	}

	var optNames []string
	for _, cmdData := range cmdsData {
		if cmdData.FlagName == domain.Flag(flag) {
			for _, opt := range cmdData.Opts {
				optNames = append(optNames, opt.Name)
			}
			break
		}
	}

	if len(optArgs) > 0 {
		for _, optArg := range optArgs {
			if domain.IsOptName(optArg) && !slices.Contains(optNames, optArg) {
				return false
			}
		}
	}
	return true
}

func hasOnlyOneFlagArgumentAtStartingPosition(flag string, args []string, allCmdData []domain.Argument) bool {
	var validFlags []domain.Flag

	for _, item := range allCmdData {
		validFlags = append(validFlags, item.FlagName)
	}

	if !slices.Contains(validFlags, domain.Flag(flag)) {
		return false
	}

	for i := 1; i < len(args); i++ {
		arg := domain.Flag(args[i])
		if slices.Contains(validFlags, arg) {
			return false
		}
	}

	return true
}
