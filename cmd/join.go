package cmd

import (
	"fmt"

	"../framework"
)

func Join(ctx *framework.Context) {
	for _, id := range *ctx.EnrolledUsers {
		if ctx.User.ID == id {
			return
		}
	}

	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers), ctx.User.ID)
	content := fmt.Sprintf("Welcome to the club, %v!\n", ctx.User.Username)
	ctx.Reply(content)
}
