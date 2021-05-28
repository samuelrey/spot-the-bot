package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

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
	content := fmt.Sprintf("Welcome to the club, %v!\n", ctx.User.Username)
	ctx.Reply(content)
}
