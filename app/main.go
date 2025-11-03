package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

type CommandHandler func(args []string) (exit bool)

func main() {
	handlers := map[string]CommandHandler{
		"exit": handleExit,
		"echo": handleEcho,
	}

	// REPL
	for {
		fmt.Fprint(os.Stdout, "$ ")
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		if handler, ok := handlers[cmd]; ok {
			if handler(args) {
				break
			}
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
