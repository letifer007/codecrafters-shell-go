package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		
		// Since the string returned by ReadString('\n') includes a trailing newline
		command = command[:len(command) - 1]

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		if command == "exit 0" {
			os.Exit(0)
		}

		fmt.Println(command + ": command not found")
	}
}
