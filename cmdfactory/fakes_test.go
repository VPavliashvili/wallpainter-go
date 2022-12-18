package cmdfactory_test

import "github.com/VPavliashvili/slideshow-go/domain"

func getFakeArgument(flag string, opts []domain.Opt) *domain.Argument {
	return &domain.Argument{
		FlagName: domain.Flag(flag),
		Opts:     opts,
	}
}

type fakeParser struct{}

func (f fakeParser) Parse(args []string) (*domain.Argument, error) {
	result := getFakeArgument(args[0], []domain.Opt{})
	switch args[0] {
	case "flag3":
		result.Opts = []domain.Opt{{
			Name:  "o2",
			Value: "v2",
		}}
	case "flag2":
		result.Opts = []domain.Opt{{
			Name:  "opt1",
			Value: "val1",
		}}
	case "flag1":
		result.Opts = []domain.Opt{{
			Name: "d",
			Value: "k",
		}}
	}
	return result, nil
}

type fakeCommand struct {
	flagName string
	opts     []domain.Opt
}

func (f fakeCommand) Execute() error { return nil }

func (f fakeCommand) GetArgument() domain.Argument {
	return domain.Argument{
		FlagName: domain.Flag(f.flagName),
		Opts:     f.opts,
	}
}

func (f fakeCommand) Name() string { return f.flagName }

func (f *fakeCommand) SetArgument(arg domain.Argument) {
	f.flagName = string(arg.FlagName)
	f.opts = arg.Opts
}

type fakeProvider struct {
	fakeCmds []domain.Command
}

func (f fakeProvider) Get() []domain.Command {
	return f.fakeCmds
}

var fakeAvailableCommands fakeProvider = fakeProvider{
	fakeCmds: []domain.Command{
		&fakeCommand{
			flagName: "flag1",
		},
		&fakeCommand{
			flagName: "flag2",
			opts: []domain.Opt{
				{
					Name:  "opt1",
					Value: "val1",
				},
			},
		},
		&fakeCommand{
			flagName: "flag3",
		},
	},
}
