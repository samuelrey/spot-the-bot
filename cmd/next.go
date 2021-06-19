package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const (
	StrSkipUser = "%s, see you next time around!"
	StrNextUser = "%s, you're up next!"
)

func Next(ctx *framework.CommandContext) {
	if len(*ctx.EnrolledUsers) <= 1 {
		return
	}

	skipUser := (*ctx.EnrolledUsers)[0]
	if ctx.Actor.ID != skipUser.ID {
		return
	}

	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers)[1:], skipUser)
	content := fmt.Sprintf(StrSkipUser, ctx.Actor)
	ctx.Reply(content)

	nextUser := (*ctx.EnrolledUsers)[0]

	content = fmt.Sprintf(StrNextUser, nextUser)
	ctx.Reply(content)
}
