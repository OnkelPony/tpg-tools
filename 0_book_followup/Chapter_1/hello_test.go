package hello

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	hello.PrintTo(fakeTerminal)
	want := "Hello, world"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
