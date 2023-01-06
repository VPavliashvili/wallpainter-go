package cmdfactory_test

import (
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/flags"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func getFakeArgument(flag string, opts []opts.Opt) *domain.CmdArgument {
	return &domain.CmdArgument{
		Flag: flags.ToFlag(flag),
		Opts: opts,
	}
}

type fakeParser struct{}

func (f fakeParser) Parse(args []string) (*domain.CmdArgument, error) {
	result := getFakeArgument(args[0], []opts.Opt{})
	switch args[0] {
	case "flag3":
		result.Opts = []opts.Opt{{
			Name:  "o2",
			Value: "v2",
		}}
	case "flag2":
		result.Opts = []opts.Opt{{
			Name:  "opt1",
			Value: "val1",
		}}
	case "flag1":
		result.Opts = []opts.Opt{{
			Name:  "d",
			Value: "k",
		}}
	}
	return result, nil
}

type fakeCommand struct {
	flagName string
	opts     []opts.Opt
}

func (f fakeCommand) Execute() error { return nil }

func (f fakeCommand) Name() string { return f.flagName }

func (f *fakeCommand) SetArgument(arg domain.CmdArgument) {
	f.flagName = string(arg.Flag)
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
			opts: []opts.Opt{
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
