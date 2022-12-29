package domain

type Command interface {
	SetArgument(Argument)
	Execute() error
    Name() string
}

type AvailableCommandsProvider interface {
	Get() []Command
}
