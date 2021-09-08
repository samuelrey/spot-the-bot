package cmd

import (
	"fmt"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *CommandContext) {
	if ctx.UserQueue.Contains(ctx.Actor) {
		return
	}

	ctx.UserQueue.Push(ctx.Actor)
	content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
	ctx.Reply(content)
}
