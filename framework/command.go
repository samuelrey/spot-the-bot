package framework

type (
	command        func(*CommandContext)
	commandMap     map[string]command
	CommandRegistry struct {
		commandMap
	}
)

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{make(commandMap)}
}

func (cr CommandRegistry) Get(name string) (*command, bool) {
	cmd, found := cr.commandMap[name]
	return &cmd, found
}

func (cr CommandRegistry) Register(name string, cmd command) {
	cr.commandMap[name] = cmd
}

type CommandContext struct {
	Messager
	PlaylistCreator
	PlaylistName  string
	EnrolledUsers *[]MessageUser
	Actor         MessageUser
}
