package cmd

import (
	"fmt"

	"../framework"
)

func Join(ctx *framework.Context) {
	enrolled := false
	for _, id := range *ctx.UserIDs {
		if id == ctx.User.ID {
			enrolled = true
			break
		}
	}
	if !enrolled {
		*ctx.UserIDs = append(*ctx.UserIDs, ctx.User.ID)
		fmt.Println("new user: ", ctx.User.ID)
	}
}
