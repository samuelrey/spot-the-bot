package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

func List(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) == 0 {
		ctx.Reply("Nobody's joined yet, join with `!join`")
		return
	}

	usernames := make([]string, len(*ctx.EnrolledUsers))
	for i, user := range *ctx.EnrolledUsers {
		usernames[i] = user.Username
	}
	content := fmt.Sprintf("Here's the queue:\n%v\n", usernames)
	ctx.Reply(content)
}
