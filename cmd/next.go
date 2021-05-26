package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

func Next(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) <= 1 {
		return
	}

	skipUserID := (*ctx.EnrolledUsers)[0]
	if ctx.User.ID != skipUserID {
		return
	}

	(*ctx.EnrolledUsers) = append((*ctx.EnrolledUsers)[1:], skipUserID)
	content := fmt.Sprintf("%v, see you next time around!", ctx.User.Username)
	ctx.Reply(content)

	nextUserID := (*ctx.EnrolledUsers)[0]
	nextUser, err := ctx.Discord.User(nextUserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	content = fmt.Sprintf("%v, you're up next!", nextUser.Username)
	ctx.Reply(content)
}
