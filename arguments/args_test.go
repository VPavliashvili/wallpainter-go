package arguments

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

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
		unaryArgument{
			name: setValue("-h"),
		},
		binaryArgument{
			name:  setValue("-p"),
			value: setValue(picturesDir),
		},
	}

	got, err := GetArguments(args)
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

func TestCreateArgument_ForUnary_Binary(t *testing.T) {
	home, _ := os.UserHomeDir()
	picturesDir := fmt.Sprintf("%v/Pictures/", home)

	fake := []struct {
		key   string
		value string
		want  Argument
	}{
		{
			key:   "--path",
			value: picturesDir,
			want: binaryArgument{
				name:  setValue("--path"),
				value: &picturesDir,
				desc: new(string),
			},
		},
		{
			key:   "-h",
			value: "",
			want:  unaryArgument{
				name: setValue("-h"),
				desc: new(string),
			},
		},
	}

    for _, item := range fake {
        get, _ := createArgument(item.key, item.value)
        want := item.want
        if get.GetName() != want.GetName() || get.Value() != want.Value() {
            t.Errorf("error in createArgument\nget\n%vwant\n%v\ncase\n%v", get, want, item)
        }
    }
}

func TestCreateValidArgument(t *testing.T) {
	home, _ := os.UserHomeDir()
	picturesDir := fmt.Sprintf("%v/Pictures/", home)

	var testItem = func(key string, value string, want Argument) {
		got, err := createArgument(key, value)
		if err != nil {
			t.Errorf("error is not expected here. got %v, want %v", got, want)
		}
		assertEqual(t, want.Value(), got.Value())
		assertEqual(t, want.GetName(), got.GetName())
	}

	data := []struct {
		key   string
		value string
		want  Argument
	}{
		{
			key:   "-h",
			value: "",
			want: unaryArgument{
				name: setValue("-h"),
			},
		},
		{
			key:   "--path",
			value: picturesDir,
			want: binaryArgument{
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
