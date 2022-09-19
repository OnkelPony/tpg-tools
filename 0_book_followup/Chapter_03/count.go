package count

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type option func(*counter) error

type counter struct {
	input  io.Reader
	output io.Writer
}

func (c counter) Lines() int {
	var result int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		result++
	}
	return result
}

func Lines() int {
	c, err := NewCounter()
	if err != nil {
		panic("internal error")
	}
	return c.Lines()
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}
