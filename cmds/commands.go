package cmds

import "github.com/VPavliashvili/slideshow-go/domain"

type availableCommands struct {}

func (ac availableCommands) Get() []domain.Command{
    return available
}

func GetProvider() domain.AvailableCommandsProvider {
    return availableCommands{}
}

var available = []domain.Command{

}
