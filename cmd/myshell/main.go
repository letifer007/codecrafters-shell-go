package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	const (
		EXIT = "exit"
		ECHO = "echo"
		TYPE = "type"
		BUILTIN = " is a shell builtin"
		PWD = "pwd"
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
			case ECHO, TYPE, EXIT, PWD:
				fmt.Fprintln(os.Stdout, args + " is a shell builtin")
			default:
				found := false
				for _, path := range paths {
					executable := path + "/" + args
					if _, err := os.Stat(executable); err == nil {
						fmt.Fprintln(os.Stdout, args + " is " + executable)
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
		} else if cmd == PWD {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error printing working directory:", err)
			}
			fmt.Fprintln(os.Stdout, cwd)
		} else {
			found := false
			for _, path := range paths {
				executable := path + "/" + cmd
				if _, err := os.Stat(executable); err == nil {
					execCmd := exec.Command(cmd, args)
					output, err := execCmd.CombinedOutput()
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error executing the program:", err)
						break
					} else {
						fmt.Fprint(os.Stdout, string(output))
						found = true
						break
					}
				}
			} 
			if !found {
				fmt.Fprintln(os.Stderr, input + ": command not found")
			}
		}
	}
}
