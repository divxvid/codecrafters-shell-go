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
		if len(c.Args) != 1 {
			return fmt.Errorf("Invalid Count of Args")
		}

		value, err := strconv.Atoi(c.Args[0])
		if err != nil {
			return err
		}

		os.Exit(value)
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
