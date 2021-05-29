package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *framework.Context) {
	for _, user := range *ctx.EnrolledUsers {
		if ctx.Actor.ID == user.ID {
			return
		}
	}

	user := framework.User{
		ID:       ctx.Actor.ID,
		Username: ctx.Actor.Username,
	}
	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers), user)
	content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
	ctx.Reply(content)
}
