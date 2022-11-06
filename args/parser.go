package args

import (
	"github.com/VPavliashvili/slideshow-go/cmdconf"
	"golang.org/x/exp/slices"
)

func CreateParser(data cmdconf.AllCommandData) Parser {
	return parser{
		allCommandsData: data,
	}
}

type Parser interface {
	Parse([]string) (*cmdconf.Data, error)
}

type parser struct {
	allCommandsData cmdconf.AllCommandData
}

func (b parser) Parse(args []string) (*cmdconf.Data, error) {
	cmdsData := b.allCommandsData.GetAllCommandData()
	flag := args[0]
	optArgs := args[1:]

	if !hasOnlyOneFlagArgumentAtStartingPosition(flag, args, cmdsData) ||
		!checkAllEnteredFlagOptionArgumentsAreValid(flag, optArgs, cmdsData) {
		return nil, invalidInputError{input: args}
	}

	result := getCmddata(flag, optArgs)
	return &result, nil
}

func getCmddata(flag string, optArgs []string) cmdconf.Data {
	var result cmdconf.Data
	result.FlagName = cmdconf.Flag(flag)
	pairs := make(map[string]string)

	for i := 0; i < len(optArgs); i++ {
		arg := optArgs[i]
		if i < len(optArgs)-1 {
			next := optArgs[i+1]
			if !cmdconf.IsOptName(next) {
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
		result.Opts = append(result.Opts, cmdconf.Opt{
			Name:  opt,
			Value: val,
		})
	}

	return result
}

func checkAllEnteredFlagOptionArgumentsAreValid(flag string, optArgs []string, cmdsData []cmdconf.Data) bool {
	if len(optArgs) > 0 && !cmdconf.IsOptName(optArgs[0]) {
		return false
	}

	var optNames []string
	for _, cmdData := range cmdsData {
		if cmdData.FlagName == cmdconf.Flag(flag) {
			for _, opt := range cmdData.Opts {
				optNames = append(optNames, opt.Name)
			}
			break
		}
	}

	if len(optArgs) > 0 {
		for _, optArg := range optArgs {
			if cmdconf.IsOptName(optArg) && !slices.Contains(optNames, optArg) {
				return false
			}
		}
	}
	return true
}

func hasOnlyOneFlagArgumentAtStartingPosition(flag string, args []string, allCmdData []cmdconf.Data) bool {
	var validFlags []cmdconf.Flag

	for _, item := range allCmdData {
		validFlags = append(validFlags, item.FlagName)
	}

	if !slices.Contains(validFlags, cmdconf.Flag(flag)) {
		return false
	}

	for i := 1; i < len(args); i++ {
		arg := cmdconf.Flag(args[i])
		if slices.Contains(validFlags, arg) {
			return false
		}
	}

	return true
}
