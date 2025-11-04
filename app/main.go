package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandHandlers map[string]func([]string)

// Initialize with built in function
func init() {
	commandHandlers = map[string]func([]string){
		"echo": echoHandler,
		"exit": exitHandler,
		"type": typeHandler,
		"pwd":  pwdHandler,
		"cd":   cdHandler,
	}
}

func main() {
	for {
		// Input
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Output
		commands := strings.Split(input, " ")
		handler, ok := commandHandlers[strings.TrimSpace(commands[0])]

		if ok {
			handler(commands) // Builtin handler
		} else if anotherProgramHandler(commands) { // Try to run program in system
			continue
		} else { // Program not found
			fmt.Fprintln(os.Stdout, input[:len(input)-1]+": command not found")
			continue
		}
	}
}
