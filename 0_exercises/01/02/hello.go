package hello

import (
	"bytes"
	"fmt"
)

func PrintTo(terminal *bytes.Buffer, keyboard *bytes.Buffer) {
	fmt.Fprintf(terminal, "Hello, [%s]", keyboard.String())
}
