package cmd

import (
	"fmt"

	"../framework"
)

func List(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) == 0 {
		ctx.Reply("Nobody's joined yet, join with `!join`")
		return
	}

	usernames := getEnrolledUsernames(ctx)
	content := fmt.Sprintf("Here's the queue:\n%v\n", usernames)
	ctx.Reply(content)
}

func getEnrolledUsernames(ctx *framework.Context) []string {
	usernames := make([]string, len(*ctx.EnrolledUsers))
	for i, id := range *ctx.EnrolledUsers {
		u, err := ctx.Discord.User(id)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		usernames[i] = u.Username
	}
	return usernames
}
