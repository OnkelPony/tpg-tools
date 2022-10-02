package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func NewSession(i io.Reader, o io.Writer, e io.Writer) *Session {
	return &Session{Stdin: i, Stdout: o, Stderr: e}
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprint(s.Stdout, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.Stdout, "\nBe seeing you!")
			break
		}
		cmd, err := CmdFromString(line)
		if err != nil {
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, line)
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s", out)
	}
}

func RunCLI() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}
