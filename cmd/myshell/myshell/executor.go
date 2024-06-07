package myshell

import (
	"fmt"
	"io"
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
	if !found {
		fmt.Fprintf(w, "%s: command not found\n", command.CommandName)
		return
	}

	err := f(command, w)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v\n", err)
	}
}
