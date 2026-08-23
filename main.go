package main

import (
	"os"

	"stc-redis/internal/commands"
	"stc-redis/internal/repl"
)

func main() {
	repl.Run(os.Stdin, os.Stdout, commands.Dispatch)
}
