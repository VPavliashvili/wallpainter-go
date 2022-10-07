package args

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestCreateValidArgument(t *testing.T) {
	home, _ := os.UserHomeDir()
	pictures := fmt.Sprintf("%v/Pictures/", home)

	pairs := []struct {
		key   string
		value string
	}{
		{key: "-h", value: "true"},
		{key: "--path", value: pictures},
	}

	hpair := pairs[0]
	got, err := createArgument(hpair.key, hpair.value)
	want := argument{
		name:  setValue("-h"),
		value: setValue("true"),
	}
	if err != nil {
		t.Errorf("error is not expected here. got %v, want %v", got, want)
	}
	assertEqual(t, got.Value(), *want.value)
	assertEqual(t, got.GetName(), *want.name)

	ppair := pairs[1]
	got, err = createArgument(ppair.key, ppair.value)
	want = argument{
		name:  setValue("--path"),
		value: setValue(pictures),
	}
	if err != nil {
		t.Errorf("error is not expected here. got %v, want %v", got, want)
	}
	assertEqual(t, got.Value(), *want.value)
	assertEqual(t, got.GetName(), *want.name)

}

func TestCreateInvalidArgument(t *testing.T) {
	invalids := []struct {
		key   string
		value string
	}{
		{key: "--nonexistent", value: "does not really matter"},
		{key: "-p", value: "10"},
	}

	withInvalidKey := invalids[0]
	resp, got := createArgument(withInvalidKey.key, withInvalidKey.value)
	if resp != nil {
		t.Errorf("response should be nil when invalid key is passed to createArgument function, returned %v", resp)
	}
	want := argError{
		name:  withInvalidKey.key,
		value: withInvalidKey.value,
	}
	if !errors.Is(got, want) {
		t.Errorf("innacurate error. got %v, want %v", got, want)
	}

	withInvalidValue := invalids[1]
	resp, got = createArgument(withInvalidValue.key, withInvalidValue.value)
	if resp != nil {
		t.Errorf("response should be nil when invalid value is passed to createArgument function, returned %v", resp)
	}
	want = argError{
		name:  withInvalidValue.key,
		value: withInvalidValue.value,
	}
	if !errors.Is(got, want) {
		t.Errorf("innacurate error. got %v, want %v", got, want)
	}

}
