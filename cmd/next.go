package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

func Next(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) <= 1 {
		return
	}

	skipUser := (*ctx.EnrolledUsers)[0]
	if ctx.User.ID != skipUser.ID {
		return
	}

	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers)[1:], skipUser)
	content := fmt.Sprintf("%v, see you next time around!", ctx.User.Username)
	ctx.Reply(content)

	nextUser := (*ctx.EnrolledUsers)[0]

	content = fmt.Sprintf("%v, you're up next!", nextUser.Username)
	ctx.Reply(content)
}
