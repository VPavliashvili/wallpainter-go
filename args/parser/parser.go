package parser

import (
	"errors"

	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"golang.org/x/exp/slices"
)

func Create(data domain.AvailableArgumentsProvider) Parser {
	return concreteParser{
		allArgumentsData: data,
	}
}

type Parser interface {
	Parse([]string) (*domain.Argument, error)
}

type concreteParser struct {
	allArgumentsData domain.AvailableArgumentsProvider
}

func (b concreteParser) Parse(args []string) (*domain.Argument, error) {
	if len(args) == 0 {
		return nil, errors.New("args array should not be empty")
	}

	argsData := b.allArgumentsData.Get()
	flag := args[0]
	optArgs := args[1:]

	if flag == flags.SetWallpaper {
		return &domain.Argument{
			Flag: flags.ToFlag(flag),
			Opts: getSetWallpaperCommandOpts(optArgs),
		}, nil
	}

	if !hasOnlyOneFlagArgumentAtStartingPosition(flag, args, argsData) ||
		!checkAllEnteredFlagOptionArgumentsAreValid(flag, optArgs, argsData) {
		return nil, invalidInputError{input: args}
	}

	result := getCmddata(flag, optArgs)
	return &result, nil
}

func getSetWallpaperCommandOpts(optArgs []string) []domain.Opt {
	var result []domain.Opt
	var usedIndexes []int
	for i := 0; i < len(optArgs); i++ {
		arg := optArgs[i]
		if domain.IsOptName(arg) {
			next := domain.Opt{Name: arg, Value: optArgs[i+1]}
			usedIndexes = append(usedIndexes, i+1)
			result = append(result, next)
		} else {
			if !slices.Contains(usedIndexes, i) {
				next := domain.Opt{Name: "", Value: arg}
				result = append(result, next)
			}
		}
	}

	return result
}

func getCmddata(flag string, optArgs []string) domain.Argument {
	var result domain.Argument
	result.Flag = flags.ToFlag(flag)
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
		if cmdData.Flag == flags.ToFlag(flag) {
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
	var validFlags []flags.Flag

	for _, item := range allCmdData {
		validFlags = append(validFlags, item.Flag)
	}

	if !slices.Contains(validFlags, flags.ToFlag(flag)) {
		return false
	}

	for i := 1; i < len(args); i++ {
		arg := flags.ToFlag(args[i])
		if slices.Contains(validFlags, arg) {
			return false
		}
	}

	return true
}
