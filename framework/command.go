package framework

type (
	Command        func(*CommandContext)
	CommandMap     map[string]Command
	CommandRegistry struct {
		CommandMap
	}
)

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{make(CommandMap)}
}

func (cr CommandRegistry) Get(name string) (*Command, bool) {
	command, found := cr.CommandMap[name]
	return &command, found
}

func (cr CommandRegistry) Register(name string, command Command) {
	cr.CommandMap[name] = command
}

type CommandContext struct {
	Messager
	PlaylistCreator
	PlaylistName  string
	EnrolledUsers *[]MessageUser
	Actor         MessageUser
}
