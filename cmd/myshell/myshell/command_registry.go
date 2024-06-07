package myshell

import (
	"os"
	"path/filepath"
	"strings"
)

type CommandRegistry struct {
	registeredCommands map[string]ExecutorFunc
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		registeredCommands: make(map[string]ExecutorFunc),
	}
}

func (cr *CommandRegistry) Register(commandName string, f ExecutorFunc) {
	cr.registeredCommands[commandName] = f
}

func (cr *CommandRegistry) GetExecutor(commandName string) (ExecutorFunc, bool) {
	f, found := cr.registeredCommands[commandName]
	return f, found
}

func (cr *CommandRegistry) GetCommandPath(commandName string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, ":")

	for _, path := range paths {
		path = strings.TrimSpace(path)
		entries, err := os.ReadDir(path)

		if err != nil {
			//there was some issue with the path variable path
			//we can't do much so just continue
			continue
		}
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			permission := info.Mode()
			if info.IsDir() || permission&0111 == 0 {
				//it is not executable by anyone
				continue
			}
			fname := entry.Name()
			fullPath := filepath.Join(path, fname)
			if fname == commandName {
				return fullPath, true
			}
		}
	}

	return "", false
}
