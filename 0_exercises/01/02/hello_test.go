package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	fakeKeyboard := &bytes.Buffer{}
	fakeKeyboard.WriteString("Jirka")
	hello.PrintTo(fakeTerminal, fakeKeyboard)
	want := "Hello, [Jirka]"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
