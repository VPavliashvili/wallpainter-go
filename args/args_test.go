package args

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type fakeTrimmer struct{
}
func (fakeTrimmer) Trim(s []string) ([]string, error){
    return s, nil
}

func TestGetArguments(t *testing.T) {

	var equal = func(got Argument, want Argument) {
		if got.GetName() != want.GetName() || got.Value() != want.Value() {
			t.Errorf("these two args should have been equal. got %v, want %v", got, want)
		}
	}

	home, _ := os.UserHomeDir()
	picturesDir := fmt.Sprintf("%v/Pictures/", home)
	args := []string{
		"-h", "-p", picturesDir,
	}
	want := []Argument{
		argument{
			name:  setValue("-h"),
			value: setValue("true"),
		},
		argument{
			name:  setValue("-p"),
			value: setValue(picturesDir),
		},
	}

	got, err := GetArguments(args, fakeTrimmer{})
	if err != nil {
		t.Errorf("error is not expected here. got %v, want %v", got, want)
	}
	if len(got) != 2 {
		t.Errorf("wrong number of arguments have returned. got %v, want %v", len(got), 2)
	}

    var count int
	for _, arg := range got {
		for _, wnt := range want {
			if arg.GetName() == wnt.GetName() {
				equal(arg, wnt)
                count++
			}
		}
	}
    if count != len(got) {
        t.Errorf("not two of expected args are equal. got %v, want %v", count, len(got))
    }
}

func TestCreateValidArgument(t *testing.T) {
	home, _ := os.UserHomeDir()
	picturesDir := fmt.Sprintf("%v/Pictures/", home)

	var testItem = func(key string, value string, want argument) {
		got, err := createArgument(key, value)
		if err != nil {
			t.Errorf("error is not expected here. got %v, want %v", got, want)
		}
		assertEqual(t, got.Value(), *want.value)
		assertEqual(t, got.GetName(), *want.name)
	}

	data := []struct {
		key   string
		value string
		want  argument
	}{
		{
			key:   "-h",
			value: "true",
			want: argument{
				name:  setValue("-h"),
				value: setValue("true"),
			},
		},
		{
			key:   "--path",
			value: picturesDir,
			want: argument{
				name:  setValue("--path"),
				value: setValue(picturesDir),
			},
		},
	}

	for _, item := range data {
		testItem(item.key, item.value, item.want)
	}

}

func TestCreateInvalidArgument(t *testing.T) {
	var validateInvalid = func(key string, value string) {
		want := argError{
			name:  key,
			value: value,
		}
		resp, got := createArgument(key, value)
		if resp != nil {
			t.Errorf("response should be nil when invalid value is passed to createArgument function, returned %v", resp)
		}
		if !errors.Is(got, want) {
			t.Errorf("innacurate error. got %v, want %v", got, want)
		}
	}

	invalids := []struct {
		key   string
		value string
	}{
		{key: "--nonexistent", value: "does not really matter"},
		{key: "-p", value: "10"},
	}

	for _, item := range invalids {
		validateInvalid(item.key, item.value)
	}
}
