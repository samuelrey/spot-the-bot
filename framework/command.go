package framework

import "fmt"

type (
	command        struct {
		cmd func(*CommandContext)
		helpMsg string
	}
	commandMap     map[string]command
	CommandRegistry struct {
		commandMap
	}
)

func (c command) RunWithContext(ctx *CommandContext) {
	c.cmd(ctx)
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{make(commandMap)}
}

func (cr CommandRegistry) Get(name string) (*command, bool) {
	cmd, found := cr.commandMap[name]
	return &cmd, found
}

func (cr CommandRegistry) Register(name string, cmd func(*CommandContext), helpMsg string) {
	cr.commandMap[name] = command{cmd: cmd, helpMsg: helpMsg}
}

func (cr CommandRegistry) FormatHelpMessage() string {
	var helpMessage string

	for name, cmd := range cr.commandMap {
		helpMessage = helpMessage + fmt.Sprintf("%s: %s\n", name, cmd.helpMsg)
	}

	return helpMessage
}

type CommandContext struct {
	Messager
	PlaylistCreator
	PlaylistName string
	UserQueue    *UserQueue
	Actor        MessageUser
}
