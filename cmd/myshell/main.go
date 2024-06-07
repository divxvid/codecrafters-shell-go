package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Wait for user input
	executor := NewExecutor()
	executor.Register("exit", func(c *Command, w io.Writer) error {
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
	executor.Register("echo", func(c *Command, w io.Writer) error {
		fmt.Fprintf(w, "%s\n", c.Rest)
		return nil
	})
	executor.Register("type", func(c *Command, w io.Writer) error {
		c.ParseArgs()
		if len(c.Args) != 1 {
			return fmt.Errorf("Invalid Count of Args")
		}

		cName := c.Args[0]

		switch executor.Type(cName) {
		case BUILTIN:
			fmt.Fprintf(w, "%s is a shell builtin\n", cName)
		default:
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

		command := ParseCommand(line)
		executor.Execute(command, os.Stdout)
	}
}
