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
		paths := strings.Split(os.Getenv("PATH"), ":")
		
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		// Since the string returned by ReadString('\n') includes a trailing newline
		input = strings.TrimSpace(input)

		input_split := strings.SplitN(input, " ", 2)
		cmd := input_split[0]
		args := ""
		if len(input_split) > 1 {
			args = input_split[1]
		}

		if cmd == EXIT {
			os.Exit(0)
		} else if cmd == TYPE {
			switch args {
			case ECHO, TYPE, EXIT:
				fmt.Fprintln(os.Stdout, args + " is a shell builtin")
			default:
				found := false
				for _, path := range paths {
					exec := path + "/" + args
					if _, err := os.Stat(exec); err == nil {
						fmt.Fprintln(os.Stdout, args + " is " + exec)
						found = true
						break
					}
				}
				if !found {
					fmt.Fprintln(os.Stderr, args + ": not found")
				}
			}
		} else if cmd == ECHO {
			fmt.Fprintln(os.Stdout, args)
		} else {
			fmt.Fprintln(os.Stderr, input + ": command not found")
		}
	}
}
