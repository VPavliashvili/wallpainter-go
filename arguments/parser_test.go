package arguments

import (
	"errors"
	"reflect"
	"testing"
)

func errorf[T comparable](t *testing.T, want T, got T) {
	t.Errorf("want %v, got %v", want, got)
}

func assertEqual[T comparable](t *testing.T, want T, got T) {
	if want != got {
		errorf(t, want, got)
	}
}

func TestGetArgsFromConsole(t *testing.T) {
	allArgs := []struct {
		args []string
		want map[string]string
	}{
		{
			args: []string{"-h", "--path", "~/Pictures"},
			want: map[string]string{"-h": "", "--path": "~/Pictures"},
		},
		{
			args: []string{"-r", "-p", "~/Pictures", "-t", "10"},
			want: map[string]string{"-r": "", "-p": "~/Pictures", "-t": "10"},
		},
		{
			args: []string{"-p", "~/Pictures", "-t", "10"},
			want: map[string]string{"-p": "~/Pictures", "-t": "10"},
		},
		{
			args: []string{"-t", "10", "-r", "-p", "~/Pictures"},
			want: map[string]string{"-t": "10", "-r": "", "-p": "~/Pictures"},
		},
		{
			args: []string{"-t", "10", "-p"},
			want: map[string]string{"-t": "10", "-p": ""},
		},
	}

	for _, item := range allArgs {

		parsed, _ := getArgsFromConsole(item.args, fakeTrimmer{})
		if !reflect.DeepEqual(parsed, item.want) {
            t.Errorf("parse failed. got: %v, want: %v, input: %v", parsed, item.want, item.args)
		}
	}
}

func TestGetArgsFromConsoleWhenInvalid(t *testing.T) {
	data := []struct {
		args []string
		want error
	}{
		{
			args: []string{"-t", "10", "~/Pictures"},
			want: parseError{},
		},
		{
			args: []string{"10", "-t", "-r"},
			want: parseError{},
		},
	}

	for _, item := range data {
		resp, err := getArgsFromConsole(item.args, fakeTrimmer{})
		if errors.Is(err, item.want) {
            t.Errorf("argError should have been retuned. want: %v, got: %v, input: %v", item.want, err, item.args)
		}
		if resp != nil {
            t.Errorf("resp should have been nil, got: %v, input: %v", resp, item.args)
		}
	}
}

func TestIsCmdArg(t *testing.T) {
	data := []struct {
		arg  string
		want bool
	}{
		{arg: "-h", want: true},
		{arg: "--path", want: true},
		{arg: "~/Pictures", want: false},
	}

	for _, item := range data {
		got := isCmdArg(item.arg)
		if got != item.want {
			t.Errorf("bug in isCommangArg. want: %v, got: %v", item.want, got)
		}
	}
}
