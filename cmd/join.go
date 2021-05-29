package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const StrJoinWelcomeFmt = "Welcome to the club, %s!\n"

func Join(ctx *framework.Context) {
	for _, user := range *ctx.EnrolledUsers {
		if ctx.User.ID == user.ID {
			return
		}
	}

	user := framework.User{
		ID:       ctx.User.ID,
		Username: ctx.User.Username,
	}
	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers), user)
	content := fmt.Sprintf(StrJoinWelcomeFmt, ctx.User)
	ctx.Reply(content)
}
