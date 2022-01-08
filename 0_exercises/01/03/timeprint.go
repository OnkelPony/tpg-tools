package timeprint

import (
	"fmt"
	"io"
	"time"
)

func PrintTimeTo(w io.Writer, t time.Time) {
	fmt.Fprintf(w, "It's %d minutes past %d", t.Minute(), t.Hour())
}
