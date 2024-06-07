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
	split := strings.SplitN(line, " ", 2)

	return &Command{
		CommandName: strings.TrimSpace(split[0]),
		Rest:        strings.TrimSpace(split[1]),
	}
}
