package cmd

import (
	"fmt"
)

const StrLeaveFmt = "No hard feelings, %s!\n"

func Leave(ctx *CommandContext) {
	if ctx.UserQueue.Remove(ctx.Actor) {
		content := fmt.Sprintf(StrLeaveFmt, ctx.Actor)
		ctx.Reply(content)
	}
}
