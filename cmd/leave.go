package cmd

import (
	"fmt"

	"../framework"
)

func Leave(ctx *framework.Context) {
	if enrolled := ctx.EnrolledUsers[ctx.User.ID]; enrolled {
		ctx.EnrolledUsers[ctx.User.ID] = false

		content := fmt.Sprintf("Until next time, %v!\n", ctx.User.Username)
		ctx.Reply(content)
	}
}
