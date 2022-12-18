package domain

type Command interface {
	GetArgument() Argument
	SetArgument(Argument)
	Execute() error
    Name() string
}

type AvailableCommandsProvider interface {
	Get() []Command
}
