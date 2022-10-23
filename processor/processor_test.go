package processor

import (
	"reflect"
	"testing"
)

type fakeProvider struct {
}

func TestGetArgs(t *testing.T) {
	raw := []string{"tobeexcluded", "-h", "-r", "--path"}
	got := getArgs(raw)
	want := []string{"-h", "-r", "--path"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("error in getArgs\ngot\n%v\nwant\n%v\nraw\n%v", got, want, raw)
	}
}
