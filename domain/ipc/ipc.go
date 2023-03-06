package ipc

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"
)

type SlideshowInfo interface {
	IsRunning() bool
	GetSlideshowPictures() []string
}

func Create() SlideshowInfo {
	return slideshowInfo{}
}

type slideshowInfo struct{}

func (si slideshowInfo) IsRunning() bool {
	return false
}

func (si slideshowInfo) GetSlideshowPictures() []string {
	return nil
}

var pipefile = "mypipe"

func Read() {
	file, err := os.OpenFile(pipefile, os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %v has opened successfully", file.Name())

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print(string(line))
		}
	}
}

func Write() {

	os.Remove(pipefile)
	syscall.Mkfifo(pipefile, 0666)

	fmt.Printf("writing to %v started\n", pipefile)
	file, err := os.OpenFile(pipefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	for i := 0; true; i++ {
		text := fmt.Sprintf("hello from %vth operation\n", i+1)
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter text: ")
		// text, _ := reader.ReadString('\n')
		text = "\n" + text

		file.WriteString(text)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Printf("writing to %v finished\n", pipefile)

}
