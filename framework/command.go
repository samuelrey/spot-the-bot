package framework

type (
	Command        func(*CommandContext)
	CommandMap     map[string]Command
	CommandHandler struct {
		CommandMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CommandMap)}
}

func (ch CommandHandler) Get(name string) (*Command, bool) {
	command, found := ch.CommandMap[name]
	return &command, found
}

func (ch CommandHandler) Register(name string, command Command) {
	ch.CommandMap[name] = command
}

type CommandContext struct {
	Messager
	PlaylistCreator
	PlaylistName string
	UserQueue    UserQueue
	Actor        MessageUser
}
