package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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

		// if strings.TrimSpace(input) == "" {
		// 	continue
		// }

		commands := parseInput(input)
		// Output
		// commands := strings.Split(input, " ")
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

// func splitWithQuotes(s string) []string {
// 	var result []string
// 	var current string
// 	inQuote := false

// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '\'' {
// 			inQuote = !inQuote
// 		} else if s[i] == ' ' && !inQuote {
// 			if current != "" {
// 				result = append(result, current)
// 				current = ""
// 			}
// 		} else {
// 			current += string(s[i])
// 		}
// 	}
// 	if current != "" {
// 		result = append(result, current)
// 	}
// 	return result
// }

func parseInput(input string) []string {
	var parts []string
	var currentPart strings.Builder
	var quoteChar rune = 0

	input = strings.TrimSuffix(input, "\n")

	for _, r := range input {
		switch {
		case quoteChar != 0:
			if r == quoteChar {
				quoteChar = 0
			} else {
				currentPart.WriteRune(r)
			}
		case r == '\'' || r == '"':
			quoteChar = r

		case unicode.IsSpace(r):
			if currentPart.Len() > 0 {
				parts = append(parts, currentPart.String())
				currentPart.Reset()
			}
		default:
			currentPart.WriteRune(r)
		}
	}

	if currentPart.Len() > 0 {
		parts = append(parts, currentPart.String())
	}
	return parts
}
