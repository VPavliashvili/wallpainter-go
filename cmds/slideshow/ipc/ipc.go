package ipc

import (
	"fmt"
	"os"
	"syscall"

	"github.com/VPavliashvili/wallpainter-go/iohandler"
)

type Producer interface {
	Produce() string
}
type Receiver interface {
	Receive(string)
}

var pipefile = "mypipe"
var producer Producer
var receiver Receiver

func SetProducer(p Producer) {
	producer = p
}

func SetReceiver(r Receiver) {
	receiver = r
}

func Exists() bool {
	return iohandler.Exists(pipefile)
}

func Read() string {
	// pipe, err := os.OpenFile(pipefile, os.O_CREATE, os.ModeNamedPipe)
	pipe, err := os.OpenFile(pipefile, os.O_RDONLY, 0640)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("file %v has opened successfully\n", pipe.Name())

	data := make([]byte, 1024)
    _, _ = pipe.Read(data)

	// receiver.Receive(string(data[:]))

    return string(data[:])
}

func Write() {
	os.Remove(pipefile)
	syscall.Mkfifo(pipefile, 0666)

	fmt.Printf("writing to %v started\n", pipefile)
	pipe, err := os.OpenFile(pipefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	text := producer.Produce()
	fmt.Println(text)

	pipe.WriteString(text)
	fmt.Printf("writing to %v finished\n", pipefile)

}
