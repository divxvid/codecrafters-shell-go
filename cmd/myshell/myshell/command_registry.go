package myshell

type CommandType int

const (
	BUILTIN CommandType = iota
	NONEXISTENT
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

func (cr *CommandRegistry) GetCommandType(commandName string) CommandType {
	_, found := cr.registeredCommands[commandName]
	if found {
		return BUILTIN
	}
	return NONEXISTENT
}
