package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func checkPathInSystem(command string) (bool, string) {
	// Case: ./program /another_program
	if strings.Contains(command, "/") {
		info, err := os.Stat(command)
		if err == nil && !info.IsDir() && info.Mode().Perm()&0111 != 0 {
			return true, command
		}
		return false, ""
	}

	// Case: going through PATH
	for _, dir := range strings.Split(os.Getenv("PATH"), ":") {
		fullPath := filepath.Join(dir, command)
		info, err := os.Stat(fullPath)
		if err == nil && !info.IsDir() && info.Mode().Perm()&0111 != 0 {
			return true, fullPath
		}
	}

	return false, ""
}

func runExternal(path string, command string, args []string) {
	// Run another program inside terminal
	// Args:
	//		path : path to program
	//		command : executable file name, like python3 or ls
	//		args : array of arguments for executable command
	cmd := exec.Command(path, args...)
	cmd.Args = append([]string{command}, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
