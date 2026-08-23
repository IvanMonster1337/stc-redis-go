package commands

import (
	"fmt"

	"stc-redis/internal/resp"
)

func echo(name string, args []string) string {
	if len(args) != 1 {
		return fmt.Sprintf("-ERR wrong number of arguments for '%s' command\r\n", name)
	}

	return resp.BulkString(args[0])
}
