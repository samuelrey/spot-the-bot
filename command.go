package main

import "github.com/samuelrey/spot-the-bot/cmd"

func registerCommands(cr cmd.Registry) {
	cr.Register("join", cmd.Join, "helloWorld")
	cr.Register("leave", cmd.Leave, "helloWorld")
	cr.Register("list", cmd.List, "helloWold")
	cr.Register("next", cmd.Next, "helloWorld")
	cr.Register("create", cmd.Create, "helloWorld")
}
