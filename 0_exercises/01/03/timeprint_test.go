package timeprint_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"
	"timeprint"
)

func TestPrintFuzzyTimeToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	now := time.Now()
	timeprint.PrintTimeTo(io.Writer(fakeTerminal), now)
	want := fmt.Sprintf("It's %d minutes past %d", now.Minute(), now.Hour())
	got := fakeTerminal.String()
	//then = time.Date(0, 0, 0, 10, 8, 0, 0, nil)
	if want != got {
		t.Errorf("want = %s, got = %s", want, got)
	}
}
