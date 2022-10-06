package args

import (
	"testing"
)

func errorf[T comparable] (t *testing.T, want T, got T){
    t.Errorf("want %v, got %v", want, got)
}

func assertEqual[T comparable] (t *testing.T, want T, got T) {
    if want != got {
        errorf(t, want, got)
    }
}

func TestGetArgsFromConsole(t *testing.T) {
	args := []string{
		"-h", "--path", "~/Pictures",
	}
	parsed := getArgsFromConsole(args)

	var i int
	for key, value := range parsed {

		switch i {
		case 0:
			assertEqual(t, "--path", key)
			assertEqual(t, "~/Pictures", value)
		case 1:
			assertEqual(t, "-h", key)
			assertEqual(t, "", value)
		default:
			panic("there should only be 2 keyvalue pairs")
		}

		i++
	}
}

func TestIsCommandArg(t *testing.T) {
	args := []string{
		"-h", "--path", "~/Pictures",
	}

	for _, v := range args {
        got := isCommandArg(v)
		switch v {
        case "-h":
            assertEqual(t, true, got)
        case "--path":
            assertEqual(t, true, got)
        case "~/Pictures":
            assertEqual(t, false, got)
        default:
            panic("there should only be 3 elements")
        }

	}
}
