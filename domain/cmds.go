package domain

type Command interface {
	GetArgument() Argument
	SetArgument(Argument)
	Execute() error
}

type AvailableCommandsProvider interface {
	Get() []Command
}
