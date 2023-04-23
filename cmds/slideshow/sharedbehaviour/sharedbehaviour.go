package sharedbehaviour

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	data "github.com/VPavliashvili/wallpainter-go/domain/cmds/data/slideshow"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
	"golang.org/x/exp/slices"
)

func GetTimeOpt(options []opts.Opt) time.Duration {
	var timeoptAsString string
	contains := slices.ContainsFunc(options, func(o opts.Opt) bool {
		if o.Name == data.TimeOpt {
			timeoptAsString = o.Value
			return true
		}
		return false
	})

	if contains {
		res, _ := data.GetDurationFromOpt(timeoptAsString)
		return res
	} else {
		return data.TimeoptDefaultVal
	}

}

// this is for pictures folder
func TakeRandomElement(elemets []string, prev string) string {
    arr := []string{}

    if prev == "" {
        arr = elemets
    } else {
        for _, item := range elemets {
            if prev != item {
                arr = append(arr, item)
            }
        }
    }

	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	index := random.Intn(len(arr))
	pic := arr[index]
	return pic
}

func GetJsonStringFromFile(path string) string {
    res , err := os.ReadFile(path)
    if err != nil {
        msg := fmt.Sprintf("THIS HAPPENS WHEN %v is not created yet, it means %v is run before any other %v operation",
            data.JsonDataFileLocation, data.ListImagesOpt, data.Flag)
        panic(msg)
    }
    return string(res)
}
