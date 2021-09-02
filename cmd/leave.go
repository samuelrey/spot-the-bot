package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
)

const StrLeaveFmt = "No hard feelings, %s!\n"

func Leave(ctx *framework.CommandContext) {
	if ctx.UserQueue.Remove(ctx.Actor) {
		content := fmt.Sprintf(StrLeaveFmt, ctx.Actor)
		ctx.Reply(content)
	}
}
