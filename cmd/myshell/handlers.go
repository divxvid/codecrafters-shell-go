package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/myshell"
)

func handleCD(c *myshell.Command, w io.Writer) error {
	c.ParseArgs()
	if len(c.Args) != 1 {
		return fmt.Errorf("Incorrect number of arguments. Need 1, got %d\n", len(c.Args))
	}
	if c.Args[0] == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("Some error occured while fetching home directory. Err: %v", err)
		}
		c.Args[0] = home
	}
	err := os.Chdir(c.Args[0])
	if err != nil {
		fmt.Fprintf(w, "cd: %s: No such file or directory\n", c.Args[0])
	}
	return nil
}

func handleExit(c *myshell.Command, w io.Writer) error {
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
}

func handleType(c *myshell.Command, w io.Writer) error {
	c.ParseArgs()
	if len(c.Args) != 1 {
		return fmt.Errorf("Invalid Count of Args")
	}

	cName := c.Args[0]

	if _, found := myshell.GetCommandRegistry().GetExecutor(cName); found {
		fmt.Fprintf(w, "%s is a shell builtin\n", cName)
	} else if path, found := myshell.GetCommandRegistry().GetCommandPath(cName); found {
		fmt.Fprintf(w, "%s is %s\n", cName, path)
	} else {
		fmt.Fprintf(w, "%s not found\n", cName)
	}

	return nil
}
