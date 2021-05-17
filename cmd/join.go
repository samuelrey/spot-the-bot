package cmd

import (
	"fmt"

	"../framework"
)

func Join(ctx *framework.Context) {
	if enrolled := ctx.EnrolledUsers[ctx.User.ID]; !enrolled {
		ctx.EnrolledUsers[ctx.User.ID] = true

		content := fmt.Sprintf("Welcome to the club, %v!\n", ctx.User.Username)
		ctx.Reply(content)
	}
}
