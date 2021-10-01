package cmd

import "fmt"

type Registry struct {
	commandMap map[string]command
}

func NewRegistry() *Registry {
	return &Registry{
		commandMap: make(map[string]command),
	}
}

func (cr Registry) Get(name string) (*command, bool) {
	cmd, found := cr.commandMap[name]
	return &cmd, found
}

func (cr Registry) Register(name string, cmd func(*Context), helpMsg string) {
	cr.commandMap[name] = command{cmd: cmd, helpMsg: helpMsg}
}

func (cr Registry) FormatHelpMessage() string {
	var helpMessage string

	for name, cmd := range cr.commandMap {
		helpMessage = helpMessage + fmt.Sprintf("%s: %s\n", name, cmd.helpMsg)
	}

	return helpMessage
}
