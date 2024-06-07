package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/myshell"
)

func main() {
	// Wait for user input
	registry := myshell.NewCommandRegistry()
	executor := myshell.NewExecutor(registry)
	registry.Register("exit", func(c *myshell.Command, w io.Writer) error {
		c.ParseArgs()
		if len(c.Args) == 0 {
			return fmt.Errorf("Invalid Count of Args")
		}

		value, err := strconv.Atoi(c.Args[0])
		if err != nil {
			return err
		}

		os.Exit(value)
		return nil
	})
	registry.Register("echo", func(c *myshell.Command, w io.Writer) error {
		fmt.Fprintf(w, "%s\n", c.Rest)
		return nil
	})
	registry.Register("type", func(c *myshell.Command, w io.Writer) error {
		c.ParseArgs()
		if len(c.Args) != 1 {
			return fmt.Errorf("Invalid Count of Args")
		}

		cName := c.Args[0]

		if _, found := registry.GetExecutor(cName); found {
			fmt.Fprintf(w, "%s is a shell builtin\n", cName)
		} else if path, found := registry.GetCommandPath(cName); found {
			fmt.Fprintf(w, "%s is %s\n", cName, path)
		} else {
			fmt.Fprintf(w, "%s not found\n", cName)
		}

		return nil
	})

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
