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

func handle_exit(command string) bool {
	splited := strings.Fields(command)

	if len(splited) != 2 {
		return false
	}

	cmd := splited[0]
	status_code := splited[1]

	if cmd != "exit" {
		return false
	}

	if status_code == "0" {
		return true
	}

	return false

}

func main() {
	// REPL
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if handle_exit(command) {
			break
		}

		command = strings.TrimSpace(command)
		fmt.Printf("%s: command not found\n", command)
	}
}
