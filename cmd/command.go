package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
)

type command struct {
	cmd     func(*CommandContext)
	helpMsg string
}

func (c command) RunWithContext(ctx *CommandContext) {
	c.cmd(ctx)
}

type CommandRegistry struct {
	commandMap map[string]command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commandMap: make(map[string]command),
	}
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
	message.Messager
	playlist.PlaylistCreator
	PlaylistName string
	UserQueue    *framework.UserQueue
	Actor        message.User
}
