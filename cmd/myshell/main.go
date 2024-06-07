package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/myshell"
)

func main() {
	// Wait for user input
	registry := myshell.GetCommandRegistry()
	executor := myshell.NewExecutor(registry)
	registry.Register("exit", handleExit)
	registry.Register("echo", func(c *myshell.Command, w io.Writer) error {
		fmt.Fprintf(w, "%s\n", c.Rest)
		return nil
	})
	registry.Register("type", handleType)
	registry.Register("cd", handleCD)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input from the stdin")
			os.Exit(1)
		}

		command := myshell.ParseCommand(line)
		executor.Execute(command, os.Stdout)
	}
}
