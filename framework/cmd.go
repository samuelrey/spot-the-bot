package framework

import "fmt"

type (
	Command        func()
	CmdMap         map[string]Command
	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	cmdHandler := CommandHandler{
		cmds: map[string]Command{
			"optin": func() {
				fmt.Println("optin")
			},
			"optout": func() {
				fmt.Println("optout")
			},
		},
	}
	return &cmdHandler
}

func (cmdHandler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := cmdHandler.cmds[name]
	return &cmd, found
}
