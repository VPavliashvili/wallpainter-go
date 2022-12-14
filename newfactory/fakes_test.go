package newfactory_test

import "github.com/VPavliashvili/slideshow-go/domain"

type fakeParser struct{ fakearg string }

func (f fakeParser) Parse(args []string) (*domain.Argument, error) {
	result := &domain.Argument{
		FlagName: domain.Flag(f.fakearg),
		Opts: []domain.Opt{
			{
				Name:  "fakeopt",
				Value: "fakeval",
			},
		},
	}
	return result, nil
}

type fakeCommand struct {
	flagName string
}

func (f fakeCommand) Execute() error { return nil }

func (f fakeCommand) GetArgument() domain.Argument {
	return domain.Argument{
		FlagName: domain.Flag(f.flagName),
		Opts: []domain.Opt{
			{
				Name:  "fakeopt",
				Value: "fakeval",
			},
		},
	}
}

func (f fakeCommand) SetArgument(domain.Argument) {}

type fakeProvider struct {
	fakeCmds []domain.Command
}

func (f fakeProvider) Get() []domain.Command {
	return f.fakeCmds
}

var provider fakeProvider = fakeProvider{
	fakeCmds: []domain.Command{
		fakeCommand{
			flagName: "flag1",
		},
		fakeCommand{
			flagName: "flag2",
		},
	},
}
