package hello

import (
	"bytes"
	"fmt"
)

func PrintTo(buf *bytes.Buffer) {
	fmt.Fprintf(buf, "Hello, world")
}
