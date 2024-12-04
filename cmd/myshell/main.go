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
		TYPE = "type"
		BUILTIN = " is a shell builtin"
	)
	
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
		// Since the string returned by ReadString('\n') includes a trailing newline
		input = strings.TrimSpace(input)
		input_split := strings.SplitN(input, " ", 2)

		if strings.HasPrefix(input, TYPE) {
			if input_split[1] == EXIT {
				fmt.Fprintln(os.Stdout, EXIT+BUILTIN)
			} else if input_split[1] == ECHO {
				fmt.Fprintln(os.Stdout, ECHO+BUILTIN)
			} else if input_split[1] == TYPE {
				fmt.Fprintln(os.Stdout, TYPE+BUILTIN)
			} else {
				fmt.Fprintln(os.Stdout, input_split[1] + ": not found")
			}
			continue
		}

		if strings.HasPrefix(input, ECHO) {
			args := input_split[1]
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
