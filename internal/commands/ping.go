package commands

import "stc-redis/internal/resp"

func ping(name string, args []string) string {
	if len(args) > 0 {
		return resp.BulkString(args[0])
	}
	return "+PONG\r\n"
}
