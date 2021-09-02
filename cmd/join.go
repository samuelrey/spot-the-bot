package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *framework.CommandContext) {
	if ctx.UserQueue.Contains(ctx.Actor) {
		return
	}

	ctx.UserQueue.Push(ctx.Actor)
	content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
	ctx.Reply(content)
}
