package counting_test

import (
	"bytes"
	"counting"
	"testing"
)

const setCount = 108

func TestNext(t *testing.T) {
	t.Parallel()
	c := counting.Counter{
		Count: setCount,
	}
	for i := setCount; i < 123; i++ {
		want := i
		got := c.Next()
		if want != got {
			t.Errorf("want = %d, got = %d", want, got)
		}
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	fakeTerminal := bytes.NewBuffer(make([]byte, 0, 108))

}
