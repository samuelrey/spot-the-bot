package framework

type (
	Command        func(*Context)
	CmdMap         map[string]Command
	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (cmdHandler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := cmdHandler.cmds[name]
	return &cmd, found
}

func (cmdHandler CommandHandler) Register(name string, command Command) {
	cmdHandler.cmds[name] = command
}
