package main

import (
	"strings"
)

type Command struct {
	CommandName string
	Args        []string
	KeywordArgs map[string]string
}

func ParseCommand(line string) *Command {
	split := strings.Split(line, " ")
	args := make([]string, 0)

	for _, word := range split[1:] {
		args = append(args, strings.TrimSpace(word))
	}

	return &Command{
		CommandName: strings.TrimSpace(split[0]),
		Args:        args,
	}
}
