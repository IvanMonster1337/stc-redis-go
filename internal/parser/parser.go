package parser

import "strings"

func ParseLine(line string) []string {
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
