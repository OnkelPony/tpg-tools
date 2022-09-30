package shell_test

import (
	"github.com/google/go-cmp/cmp"
	"shell"
	"testing"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	input := "/usr/bin/ls -l main.go"
	want := []string{"/usr/bin/ls", "-l", "main.go"}
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if nil == err {
		t.Fatal("want error on empty input, got nil")
	}
}
