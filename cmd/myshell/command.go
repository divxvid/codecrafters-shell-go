package main

import (
	"strings"
)

type Command struct {
	CommandName string
	Rest        string
	Args        []string
	KeywordArgs map[string]string
}

func ParseCommand(line string) *Command {
	commandName, rest, found := strings.Cut(line, " ")
	if !found {
		rest = ""
	}

	return &Command{
		CommandName: strings.TrimSpace(commandName),
		Rest:        strings.TrimSpace(rest),
	}
}
