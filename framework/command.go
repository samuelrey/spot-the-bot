package framework

type (
	Command        func(*CommandContext)
	CmdMap         map[string]Command
	CommandHandler struct {
		CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (cmdHandler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := cmdHandler.CmdMap[name]
	return &cmd, found
}

func (cmdHandler CommandHandler) Register(name string, command Command) {
	cmdHandler.CmdMap[name] = command
}

type CommandContext struct {
	Messager
	EnrolledUsers *[]User
	Actor         User
}
