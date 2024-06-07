package myshell

import (
	"fmt"
	"io"
	"os/exec"
)

type ExecutorFunc func(*Command, io.Writer) error

type Executor struct {
	registry *CommandRegistry
}

func NewExecutor(registry *CommandRegistry) *Executor {
	return &Executor{
		registry: registry,
	}
}

func (e *Executor) Execute(command *Command, w io.Writer) {
	f, found := e.registry.GetExecutor(command.CommandName)
	if found {
		err := f(command, w)
		if err != nil {
			fmt.Fprintf(w, "An Error occured: %v\n", err)
		}
		return
	}

	path, found := e.registry.GetCommandPath(command.CommandName)
	if found {
		command.ParseArgs()
		cmd := exec.Command(path, command.Args...)
		output, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(w, "Could not execute the command: %s due to err: %v\n", path, err)
			return
		}
		w.Write(output)
		return
	}

	fmt.Fprintf(w, "%s: command not found\n", command.CommandName)
}
