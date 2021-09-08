package cmd

import (
	"fmt"
)

const (
	StrListNoUsers  = "Nobody's joined yet, join with `!join`"
	StrListUsersFmt = "Here's the queue:\n%s\n"
)

func List(ctx *CommandContext) {
	if ctx.UserQueue.Length() == 0 {
		ctx.Reply(StrListNoUsers)
		return
	}

	content := fmt.Sprintf(StrListUsersFmt, ctx.UserQueue)
	ctx.Reply(content)
}
