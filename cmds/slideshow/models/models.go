package models

import "github.com/VPavliashvili/wallpainter-go/cmds/slideshow/ipc"

type Operation interface {
	Execute() error
}

type ProducerOperation interface {
    Operation
    ipc.Producer
}

type ReceiverOperation interface {
    Operation
    ipc.Receiver
}
