package main

import (
	"fmt"
	"strings"
)

func handleExit(args []string) bool {
	if len(args) > 0 && args[0] == "0" {
		return true
	}
	return false
}

func handleEcho(args []string) bool {
	fmt.Println(strings.Join(args, " "))
	return false
}
