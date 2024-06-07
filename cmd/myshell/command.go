package main

import "strings"

type Command struct {
	commandName string
	args        []string
	keywordArgs map[string]string
}

func ParseCommand(line string) *Command {
	split := strings.Split(line, " ")

	return &Command{
		commandName: strings.TrimSpace(split[0]),
	}
}
