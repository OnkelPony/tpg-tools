package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func (c Counter) Lines() int {
	return LinesFrom(c.Input)
}

func Lines() int {
	return NewCounter().Lines()
}

func NewCounter() Counter {
	return Counter{
		Input: os.Stdin,
	}
}

func LinesFrom(r io.Reader) int {
	var result int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		result++
	}
	return result
}
