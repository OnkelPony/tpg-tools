package count_test

import (
	"count"
	"strings"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("1\n0\n8")
	want := 3
	got := count.LinesFrom(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
