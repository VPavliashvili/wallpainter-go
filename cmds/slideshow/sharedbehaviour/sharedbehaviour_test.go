package sharedbehaviour_test

import (
	"testing"

	"github.com/VPavliashvili/wallpainter-go/cmds/slideshow/sharedbehaviour"
)

func TestTakeRandomElement(t *testing.T) {
	arr := []string{"one", "two"}

	prev := ""
	for k := 0; k < 1000; k++ {
		func() {
			after := []string{}
			for range arr {
				prev = sharedbehaviour.TakeRandomElement(arr, prev)
				after = append(after, prev)
			}

			f := after[0]
			s := after[1]
			if (f == "one" && s == "one") || (f == "two" && s == "two") {
				t.Errorf("takeRandomElement error\nfirst: %v, second: %v\nafter array:%v", f, s, after)
			}
		}()
	}
}
