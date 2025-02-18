package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type option func(*counter) error

type counter struct {
	input     io.Reader
	output    io.Writer
	wordCount bool
}

func (c counter) Lines() int {
	var result int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		result++
	}
	return result
}

func (c counter) Words() int {
	var words int
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return words
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		log.Fatalln(err)
	}
	if c.wordCount {
		fmt.Println(c.Words())
	} else {
		fmt.Println(c.Lines())
	}
}

func Lines() int {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return c.Lines()
}

func Words() int {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return c.Words()
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

func FromArgs(args []string) option {
	return func(c *counter) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		wordCount := fset.Bool("w", false, "Count words instead of lines")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		args = fset.Args()
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}
