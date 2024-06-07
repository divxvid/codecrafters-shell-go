package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input from the stdin")
		os.Exit(1)
	}
	command = strings.TrimSpace(command)

	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}
