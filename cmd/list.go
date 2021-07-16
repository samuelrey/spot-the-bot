package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
)

const (
	StrListNoUsers  = "Nobody's joined yet, join with `!join`"
	StrListUsersFmt = "Here's the queue:\n%s\n"
)

func List(ctx *framework.CommandContext) {
	if ctx.UserQueue.Length() == 0 {
		ctx.Reply(StrListNoUsers)
		return
	}

	content := fmt.Sprintf(StrListUsersFmt, ctx.UserQueue)
	ctx.Reply(content)
}
