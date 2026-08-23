package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// --- resp ---

func simpleString(s string) string {
	return fmt.Sprintf("+%s\r\n", s)
}

func errorString(m string) string {
	return fmt.Sprintf("-%s\r\n", m)
}

func bulkString(s string, ok bool) string {
	if !ok {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

// --- parser ---

func parseLine(line string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false
	hasToken := false
	for _, ch := range line {
		switch {
		case ch == '"' && !inQuotes:
			inQuotes = true
			hasToken = true
		case ch == '"' && inQuotes:
			inQuotes = false
		case ch == ' ' && !inQuotes:
			if hasToken {
				args = append(args, current.String())
				current.Reset()
				hasToken = false
			}
		default:
			current.WriteRune(ch)
			hasToken = true
		}
	}
	if hasToken {
		args = append(args, current.String())
	}
	return args
}

// --- commands ---

type handler func(name string, args []string) string

var handlers = map[string]handler{
	"ECHO": echo,
	"PING": ping,
}

func dispatch(args []string) string {
	cmd := strings.ToUpper(args[0])
	if h, ok := handlers[cmd]; ok {
		return h(cmd, args[1:])
	}
	return errorString(fmt.Sprintf("ERR unknown command '%s'", cmd))
}

func ping(name string, args []string) string {
	if len(args) > 0 {
		return bulkString(args[0], true)
	}
	return simpleString("PONG")
}

func echo(name string, args []string) string {
	if len(args) != 1 {
		return errorString(fmt.Sprintf("ERR wrong number of arguments for '%s' command", name))
	}
	return bulkString(args[0], true)
}

// --- repl ---

func run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fmt.Fprint(out, dispatch(parseLine(line)))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
