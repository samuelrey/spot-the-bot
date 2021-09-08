package cmd

import (
	"fmt"
)

const (
	StrSkipUser = "%s, see you next time around!"
	StrNextUser = "%s, you're up next!"
)

func Next(ctx *CommandContext) {
	head := ctx.UserQueue.Head()
	if head == nil || ctx.Actor.ID != head.ID {
		return
	}

	ctx.UserQueue.Pop()
	ctx.UserQueue.Push(*head)
	content := fmt.Sprintf(StrSkipUser, ctx.Actor)
	ctx.Reply(content)

	head = ctx.UserQueue.Head()

	content = fmt.Sprintf(StrNextUser, head)
	ctx.Reply(content)
}
