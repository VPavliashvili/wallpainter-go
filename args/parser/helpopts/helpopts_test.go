package helpopts_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/VPavliashvili/wallpainter-go/args/parser/helpopts"
	"github.com/VPavliashvili/wallpainter-go/domain"
	"github.com/VPavliashvili/wallpainter-go/domain/opts"
)

func TestOptsParsingWhenValid(t *testing.T) {
	cases := []struct {
		opts []string
		want []opts.Opt
	}{
		{
			opts: []string{},
			want: []opts.Opt{},
		},
	}

    parser := helpopts.Create()

	for _, item := range cases {
		got, err := parser.Parse(item.opts)
		want := item.want

		if err != nil {
			t.Errorf("err should have been null\n%v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error in help opts parsing\ngot\n%v\nwant\n%v", got, want)
		}
	}
}

func TestOptsParsinWhenError(t *testing.T) {
	cases := []struct {
		opts []string
		err  error
	}{
		{
			opts: []string{"anything"},
			err: domain.InvalidOptionsError{
				OptArgs: []string{"anything"},
			},
		},
	}

    parser := helpopts.Create()

	for _, item := range cases {
		res, got := parser.Parse(item.opts)
		want := item.err

		if res != nil {
			t.Errorf("result should have been nil\n%v",res)
		}

		if !errors.Is(got, want) {
			t.Errorf("error in help opts when should throw error\ngot\n%v\nwant\n%v", got, want)
		}
	}
}
