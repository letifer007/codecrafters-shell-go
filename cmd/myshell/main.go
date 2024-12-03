package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		EXIT = "exit"
		ECHO = "echo"
	)
	
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
		// Since the string returned by ReadString('\n') includes a trailing newline
		input = strings.TrimSpace(input)

		if strings.HasPrefix(input, ECHO) {
			args := strings.TrimSpace(strings.Join(strings.Split(input, "echo"), ""))
			fmt.Fprintln(os.Stdout, args)
			continue
		}

		if strings.HasPrefix(input, EXIT) {
			os.Exit(0)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		if len(input) > 1 {
			fmt.Fprintln(os.Stderr, input + ": command not found")
		}
	}
}
