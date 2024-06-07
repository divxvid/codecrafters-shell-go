package main

import (
	"fmt"
	"io"
)

type ExecutorFunc func(*Command, io.Writer) error

type Executor struct {
	registeredCommands map[string]ExecutorFunc
}

func NewExecutor() *Executor {
	return &Executor{
		registeredCommands: make(map[string]ExecutorFunc),
	}
}

func (e *Executor) Register(commandName string, f ExecutorFunc) {
	e.registeredCommands[commandName] = f
}

func (e *Executor) Execute(command *Command, w io.Writer) {
	f, found := e.registeredCommands[command.CommandName]
	if !found {
		fmt.Fprintf(w, "%s: command not found\n", command.CommandName)
		return
	}

	err := f(command, w)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v\n", err)
	}
}
