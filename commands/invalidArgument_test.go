package commands

import "testing"

func TestString(t *testing.T) {
    fake := []struct{
        cmd invalidArgumentCommand
        want string
    } {
        {
            cmd: invalidArgumentCommand{
            	input: []string{"--invalid", "--idk"},
            },
            want: "--invalid, --idk",
        },
        {
            cmd: invalidArgumentCommand{
                input: []string{"-x"},
            },
            want: "-x",
        },
        {
            cmd: invalidArgumentCommand{input: []string{"--smth"}},
            want: "--smth",
        },
    }

    for _, item := range fake {
        got := item.cmd.String()
        want := item.want
        if got != want {
            t.Errorf("String() function error\ngot\n%v\nwant\n%v\n", got, want)
        }
    }

}
