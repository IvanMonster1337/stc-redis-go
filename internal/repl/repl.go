package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"stc-redis/internal/parser"
)

type DispatchFunc func(args []string) string

func Run(in io.Reader, out io.Writer, dispatch DispatchFunc) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fmt.Fprint(out, dispatch(parser.ParseLine(line)))
	}
}
