package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

func Leave(ctx *framework.Context) {
	found := -1
	for i, user := range *ctx.EnrolledUsers {
		if ctx.User.ID == user.ID {
			found = i
		}
	}
	if found != -1 {
		(*ctx.EnrolledUsers) = append(
			(*ctx.EnrolledUsers)[:found],
			(*ctx.EnrolledUsers)[found+1:]...,
		)
		content := fmt.Sprintf("No hard feelings, %v!\n", ctx.User.Username)
		ctx.Reply(content)
	}
}
