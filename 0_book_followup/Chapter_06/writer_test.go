package writer_test

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"writer"
)

func TestWriteToFileClobbers(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/writer_text.txt"
	err := os.WriteFile(path, []byte("hello"), 0600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 0, 8}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
