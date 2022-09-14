package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func NewPrinter() Printer {
	return Printer{
		Output: os.Stdout,
	}
}

func Print() {
	NewPrinter().Print()
}

func (p Printer) Print() {
	PrintTo(p.Output)
}

func PrintTo(w io.Writer) {
	fmt.Fprintln(w, "Hello, world")
}
