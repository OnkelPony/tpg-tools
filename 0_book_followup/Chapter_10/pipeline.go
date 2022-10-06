package pipeline

import (
	"io"
)

type Pipeline struct {
	Output io.Writer
	Error  error
}

func (p *Pipeline) Stdout() {

}

func FromString(s string) *Pipeline {
	return &Pipeline{}
}
