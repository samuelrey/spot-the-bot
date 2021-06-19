package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const StrLeaveFmt = "No hard feelings, %s!\n"

func Leave(ctx *framework.CommandContext) {
	found := -1
	for i, user := range *ctx.EnrolledUsers {
		if ctx.Actor.ID == user.ID {
			found = i
		}
	}
	if found != -1 {
		(*ctx.EnrolledUsers) = append(
			(*ctx.EnrolledUsers)[:found],
			(*ctx.EnrolledUsers)[found+1:]...,
		)
		content := fmt.Sprintf(StrLeaveFmt, ctx.Actor)
		ctx.Reply(content)
	}
}
