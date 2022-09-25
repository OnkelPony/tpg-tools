package writer_test

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
	"writer"
)

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/writer_text.txt"
	want := []byte{1, 0, 8}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode()
	if perm != 0600 {
		t.Errorf("want file mod 0600, got 0%o", perm)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

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

func TestPermsClosed(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/writer_text.txt"
	// Pre-create empty file with open perms.
	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 0, 8}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode()
	if perm != 0600 {
		t.Errorf("want file mod 0600, got 0%o", perm)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
