package commands

import (
	"fmt"
	"strings"
)

type Handler func(name string, args []string) string

var handlers = map[string]Handler{
	"ECHO": echo,
	"PING": ping,
}

func Dispatch(args []string) string {
	cmd := strings.ToUpper(args[0])
	if h, ok := handlers[cmd]; ok {
		return h(cmd, args[1:])
	}

	return fmt.Sprintf("-ERR unknown command '%s'\r\n", cmd)
}
