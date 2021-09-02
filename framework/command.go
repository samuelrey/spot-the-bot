package framework

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

type CommandContext struct {
	Messager
	PlaylistCreator
	PlaylistName string
	UserQueue    *UserQueue
	Actor        MessageUser
}
