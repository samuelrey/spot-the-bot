package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *framework.CommandContext) {
	for _, user := range *ctx.EnrolledUsers {
		if ctx.Actor.ID == user.ID {
			return
		}
	}

	user := framework.MessageUser{
		ID:       ctx.Actor.ID,
		Username: ctx.Actor.Username,
	}
	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers), user)
	content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
	ctx.Reply(content)
}
